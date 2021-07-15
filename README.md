# Cheetah

Cheetah is a tool that helps to manage a small pizza team.

## Installation

1. [Download and install](https://golang.org/doc/install) Go

2. Clone this repo locally:
      
       $ git clone https://github.com/htmfilho/cheetah.git

3. Build the code:

       $ cd path/to/cheetah
       $ go build cmd/terminal/cheetah.go

4. Install the app, making it available in any terminal:

       $ go install cmd/terminal/cheetah.go

## Usage

Type:

    $ cheetah --help
    
for a complete list of arguments. All arguments are optional, with a default value.

* `--f`: Relative or absolute path to the data file. If the current directory has a file named `team.json` then it will be automatically considered.
* `--c`: Name of the ceremony taking place. If not informed then the stand up is choosen by default because that's the most frequent ceremony.
* `--d`: Run in debug mode, showing information that is not useful for the user but for the developer.

