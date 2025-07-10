package models

type TlmCmdFileLocation int

const (
	Local TlmCmdFileLocation = iota
)

type TlmCmdFileLocationInfo struct {
	DirPath string
}

type TlmCmdConfigurationInfo struct {
	TlmApid   string
	CmdApid   string
	CompoName string
}

type TlmCmdFileConfig struct {
	Location         TlmCmdFileLocation
	CmdDBInfo        []TlmCmdConfigurationInfo
	TlmDBInfo        []TlmCmdConfigurationInfo
	CmdFileInfo      []TlmCmdFileLocationInfo
	LayoutInfo       TlmCmdFileLocation
	TlmCmdConfigInfo []TlmCmdConfigurationInfo
}
