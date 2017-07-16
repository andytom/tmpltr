package template

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/AlecAivazis/survey"
	"gopkg.in/yaml.v2"
)

const (
	// The configuration file relative to the template root directory
	configName = "config.yaml"

	// The directory holding the template file and directories relative to
	// the template root directory
	templateDir = "template"
)

// -- Template stuff --

type meta struct {
	Description string `yaml:"description"`
	Author      string `yaml:"author"`
}

type question struct {
	Name     string   `yaml:"name"`
	Type     string   `yaml:"type"`
	Message  string   `yaml:"message"`
	Help     string   `yaml:"help"`
	Required bool     `yaml:"required"`
	Default  string   `yaml:"default"`
	Options  []string `yaml:"options"`
}

// Template is our representation of a set of templates and questions.
type Template struct {
	// The Templates metadata
	Meta meta `yaml:"meta"`
	// Our internal representation of the Questions to ask
	Questions []question `yaml:"questions"`
	baseDir   string
}

// New take a directory containing a template, parses it and returns a new
// Tempalate.
func New(dir string) (Template, error) {
	t := Template{
		baseDir: dir,
	}

	configFileName := filepath.Join(t.baseDir, configName)

	file, err := os.Open(configFileName)
	if err != nil {
		return t, err
	}
	defer file.Close()

	raw, err := ioutil.ReadAll(file)
	if err != nil {
		return t, err
	}

	err = yaml.Unmarshal(raw, &t)

	return t, err
}

// GetQuestions takes the templates internal representation of the Template's
// questions and converts them into a format that can be used by survey.
func (t *Template) GetQuestions() []*survey.Question {
	var questions = []*survey.Question{}

	for _, q := range t.Questions {

		sq, err := parseQuestion(q)

		if err != nil {
			// TODO - Need to handle the error here is a better way
			continue
		}

		questions = append(questions, sq)
	}

	return questions
}

// Execute applies the template to a given directory using the data that has
// been passed in.
func (t *Template) Execute(targetRoot string, data map[string]interface{}) error {
	// Walk the template dir evaluate each path as a template and then
	// execute each file as a template
	rootDir := filepath.Join(t.baseDir, templateDir)

	walker := func(path string, info os.FileInfo, err error) error {
		// First handle any incoming error by returning it to the
		// caller rather than trying to do something fancy here.
		if err != nil {
			return err
		}

		// Build our target path from the target root and the relative
		// path to the base target dir. Then evaluate this as a
		// template to get our actual target path.
		relPath, err := filepath.Rel(rootDir, path)
		if err != nil {
			return err
		}
		rawPath := filepath.Join(targetRoot, relPath)

		targetPath, err := templatePath(rawPath, data)
		if err != nil {
			return err
		}

		fmt.Printf("Processing %q\n", targetPath)

		// If this is a dir we just create the dir and set the mode.
		if info.IsDir() {
			return os.MkdirAll(targetPath, info.Mode())
		}

		// If this is a file we process the source file as a template,
		// execute it into a new file and set the mode.
		src, err := os.Open(path)
		if err != nil {
			return err
		}
		defer src.Close()

		// Create/truncate the target file and set it's mode before
		// executing the template
		dest, err := os.Create(targetPath)
		if err != nil {
			return err
		}
		defer dest.Close()

		err = dest.Chmod(info.Mode())
		if err != nil {
			return err
		}

		return renderTemplate(src, dest, data)
	}

	return filepath.Walk(rootDir, walker)
}
