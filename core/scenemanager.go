package core

import (
	"fmt"
	"strings"

	"github.com/hajimehoshi/ebiten/v2"
)

type ISceneManager interface {
	InitScene()
	UpdateScene()
	RenderScene(screen ebiten.Image)
	AddScene(scene *Scene)
	SelectScene(sceneName string)
	AddGameObject(sceneName string, gameObj *GameObject)
}

type ScenesManager struct {
	ISceneManager
	Scenes       []*Scene
	CurrentScene *Scene
}

var SceneManager = newSceneManager()

func newSceneManager() *ScenesManager {
	defaultScene := NewScene("Default Scene")
	manager := &ScenesManager{
		Scenes:       []*Scene{defaultScene},
		CurrentScene: defaultScene,
	}
	return manager
}

func (sm *ScenesManager) InitScene() {
	if sm.CurrentScene != nil {
		sm.CurrentScene.Init()
	}
}

func (sm *ScenesManager) UpdateScene() {
	if sm.CurrentScene != nil {
		sm.CurrentScene.Update()
	}
}

func (sm *ScenesManager) RenderScene(screen *ebiten.Image) {
	if sm.CurrentScene != nil {
		sm.CurrentScene.Render(screen)
	}
}

func (sm *ScenesManager) AddScene(scene *Scene) {
	sm.Scenes = append(sm.Scenes, scene)
}

func (sm *ScenesManager) findScene(sceneName string) *Scene {
	for _, scene := range sm.Scenes {
		if strings.ToLower(scene.Name) == strings.ToLower(sceneName) {
			return scene
		}
	}
	return nil
}

func (sm *ScenesManager) SelectScene(sceneName string) {
	if scene := sm.findScene(sceneName); scene != nil {
		sm.CurrentScene = scene
	}
}

func (sm *ScenesManager) AddGameObject(sceneName string, gameObj *GameObject) {
	if scene := sm.findScene(sceneName); scene != nil {
		scene.AddGameObject(gameObj)
	} else {
		fmt.Println(fmt.Sprintf("Scene '%s' not found", sceneName))
	}
}
