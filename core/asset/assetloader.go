package assets

import (
	"fmt"
	"os"

	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type IAsset interface {
}

type AssetsLoader struct {
	Assets map[string]IAsset
}

var AssetLoader = AssetsLoader{Assets: make(map[string]IAsset)}

func (al *AssetsLoader) LoadAsset(assetName, filePath string) {
	// TODO: add supports for both img and audio
	cwd, err := os.Getwd()

	if err != nil {
		fmt.Println(err.Error())
	}

	img, _, err := ebitenutil.NewImageFromFile(cwd + filePath)

	if err != nil {
		fmt.Println(err.Error())
	}

	al.Assets[assetName] = img
}

func (al *AssetsLoader) GetAsset(assetName string) IAsset {
	asset, ok := al.Assets[assetName]
	if ok {
		return asset
	} else {
		return nil
	}
}
