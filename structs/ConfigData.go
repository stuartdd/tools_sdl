// ConfigData
package structs

type ConfigData struct {
	Name                string
	ApplicationDataPath string
	KeyBindings         map[string]string
	TextureFiles        map[string]string
	DebugKeyboard       bool
	FullScreen          bool
	NonFsWidth          int32
	NonFsHeight         int32
}
