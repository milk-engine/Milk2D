package assets

type IAsset interface {
	Load(fileName string)
}

type AssetLoader struct {
	Assets map[string]IAsset
}
