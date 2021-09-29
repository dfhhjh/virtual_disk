package internal

type Command struct {
	commandname string
}

func (cmd Command) CommandExecute() {
}

type Dir struct {
	Command
}

type Md struct {
	Command
}


type Cd struct {
	Command
}

type Touch struct {
	Command
}

type Copy struct {
	Command
}

type Del struct {
	Command
}

type Rd struct {
	Command
}

type Ren struct {
	Command
}

type Move struct {
	Command
}

type Mklink struct {
	Command
}

type Save struct {
	Command
}

type Load struct {
	Command
}

type Cls struct {
	Command
}