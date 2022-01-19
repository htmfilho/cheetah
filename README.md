# Cheetah

Cheetah is a tool that helps to manage a small pizza team.

## Installation

1. [Download and install](https://golang.org/doc/install) Go

2. Clone this repo locally:
      
       $ git clone https://github.com/htmfilho/cheetah.git

3. Build the code:

       $ cd path/to/cheetah
       $ go build cmd/terminal/cheetah.go

4. Run the app:

       $ ./cheetah [-f team.json]

6. Install the app, making it available in any terminal:

       $ go install cmd/terminal/cheetah.go

7. Copy the file `team.json.example` to a convinient place and rename it to `team.json`

8. Update the `team.json` file with information about the sprint.

## Usage

Run Cheetah in the same directory of the team.json file:

    $ cd path/to/team/file
    $ cheetah

[![asciicast](https://asciinema.org/a/NIL1nDlL5aQoim9fHz44g0eYC.svg)](https://asciinema.org/a/NIL1nDlL5aQoim9fHz44g0eYC)

Further options are described in the help:

    $ cheetah --help
    
for a complete list of arguments. All arguments are optional, with a default value.

* `--f`: Relative or absolute path to the data file. If the current directory has a file named `team.json` then it will be automatically considered.
* `--c`: Name of the ceremony taking place. If not informed then the stand up is choosen by default because that's the most frequent ceremony.
* `--d`: Run in debug mode, showing information that is not useful for the user but for the developer.

## The Team File

    {
        "cycle": "Cycle Name",
        "sprint": {
            "name":  "Sprint Name",
            "start": "2021-06-21T09:00:00-05:00",
            "end":   "2021-06-30T12:00:00-05:00"
        },
        "name": "Team Name",
        "manager": "Elvis Presley",
        "members": [{
            "name": "Celine Dion",
            "assignments": [{
                "reference": "JIRA-123",
                "summary": "Do something very interesting",
                "status": "todo"
            }, {
                "reference": "JIRA-231",
                "summary": "Develop something amazing",
                "status": "progress"
            }, {
                "summary": "It's my turn to look at server alerts.",
                "type": "weekly"
            }]
        }, {
            "name": "John Travolta",
            "assignments": [{
                "reference": "JIRA-333",
                "summary": "Fix that annoying bug",
                "status": "done"
            }, {
                "reference": "JIRA-365",
                "summary": "Add tests to complete the coverage",
                "status": "progress"
            }, {
                "reference": "JIRA-477",
                "summary": "Document that complex feature",
                "status": "todo"
            }, {
                "summary": "It's my turn to look at server alerts.",
                "type": "weekly"
            }]
        }]
    }