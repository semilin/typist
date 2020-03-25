# typist
A very simple typing test for the CLI. Written in go.

# install
Clone this repository to your computer.

Make sure Go is installed, and go build main.go

Now you can run it with ./typist

# options
There are a couple of configuration options for this program, enabled by flags in the command line. They are as follows:
* -rounds - Use this to select how many sentences you want to be tested on before it gives your final result.
* -cpm    - If you prefer characters per minute over words per minute, use this to enable CPM.
* -list  - Select which list you would like to recieve sentences from. Lists can be found in the sentences folder, and you may create your own lists if you like.