package main

type node struct {
	language *language
}

func (n *node) string() string {
	return n.language.string()
}

func (n *node) init(props *properties, env environmentInfo) {
	n.language = &language{
		env:        env,
		props:      props,
		extensions: []string{"*.js", "*.ts", "package.json", ".nvmrc"},
		commands: []*cmd{
			{
				executable: "node",
				args:       []string{"--version"},
				regex:      `(?:v(?P<version>((?P<major>[0-9]+).(?P<minor>[0-9]+).(?P<patch>[0-9]+))))`,
			},
		},
		versionURLTemplate: "[%[1]s](https://github.com/nodejs/node/blob/master/doc/changelogs/CHANGELOG_V%[2]s.md#%[1]s)",
		matchesVersionFile: n.matchesVersionFile,
	}
}

func (n *node) enabled() bool {
	return n.language.enabled()
}

func (n *node) matchesVersionFile() bool {
	fileVersion := n.language.env.getFileContent(".nvmrc")
	if len(fileVersion) == 0 {
		return true
	}
	return fileVersion == n.language.activeCommand.version.full
}
