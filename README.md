# Cheetah

Cheetah is a tool that helps to manage a small pizza team.

## Usage

Type:

    $ cheetah --help
    
for a complete list of arguments. All arguments are optional with a default value.

* `--f`: Relative or absolute path to the data file. If the current directory has a file named `team.json` then it will be automatically considered.
* `--c`: Name of the ceremony taking place. If not informed then the stand up is choosen by default because that's the most frequent ceremony.
* `--d`: Run in debug mode, showing information that is not useful for the user but for the developer.
