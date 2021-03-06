package rpm

import (
	"bufio"
	"strings"
	"log"
	"io/ioutil"
	"os"
	"os/exec"
	"github.com/MXi4oyu/DockerXScan/versionfmt"
	"github.com/MXi4oyu/DockerXScan/versionfmt/rpm"
	"github.com/MXi4oyu/DockerXScan/tarutil"
	"github.com/MXi4oyu/DockerXScan/featurefmt"
	"github.com/MXi4oyu/DockerXScan/database"
)

type lister struct{}

func init()  {
	featurefmt.RegisterLister("rpm",&lister{})
}

func (l lister) ListFeatures(files tarutil.FilesMap) (features []database.FeatureVersion,errs error) {
	f, hasFile := files["var/lib/rpm/Packages"]

	if !hasFile{
		return []database.FeatureVersion{},nil
	}
	// Create a map to store packages and ensure their uniqueness
	packagesMap:=make(map[string]database.FeatureVersion)

	var pkg database.FeatureVersion
	var err error

	// Write the required "Packages" file to disk
	tmpDir,ioerr := ioutil.TempDir(os.TempDir(), "rpm")
	defer os.RemoveAll(tmpDir)
	if ioerr!=nil{
		log.Println("could not create temporary folder for RPM detection")
		return []database.FeatureVersion{},ioerr
	}

	ioerr=ioutil.WriteFile(tmpDir+"/Packages", f, 0700)
	if ioerr!=nil{
		log.Println("could not create temporary file for RPM detection")
		return []database.FeatureVersion{},ioerr
	}

	// Extract binary package names because RHSA refers to binary package names.

	out, err := exec.Command("rpm", "--dbpath", tmpDir, "-qa", "--qf", "%{NAME} %{EPOCH}:%{VERSION}-%{RELEASE}\n").CombinedOutput()

	if err!=nil{
		log.Println("could not query RPM")
		return []database.FeatureVersion{},err
	}

	scanner := bufio.NewScanner(strings.NewReader(string(out)))

	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		if len(line) != 2 {
			// We may see warnings on some RPM versions:
			// "warning: Generating 12 missing index(es), please wait..."
			continue
		}

		// Ignore gpg-pubkey packages which are fake packages used to store GPG keys - they are not versionned properly.
		if line[0] == "gpg-pubkey" {
			continue
		}

		// Parse version
		version := strings.Replace(line[1], "(none):", "", -1)
		err := versionfmt.Valid(rpm.ParserName, version)
		if err != nil {
			log.Println("could not parse package version. skipping")
			continue
		}

		pkg= database.FeatureVersion{
			Feature: database.Feature{
				Name:line[0],
			},
			Version:version,
		}

		packagesMap[pkg.Feature.Name+"#"+pkg.Version] = pkg

	}

	packages :=make([]database.FeatureVersion,0,len(packagesMap))
	for _,pkg:=range packagesMap{
		packages=append(packages,pkg)
	}

	return packages,nil
}

func (l lister) RequiredFilenames() []string {
	return []string{"var/lib/rpm/Packages"}
}

