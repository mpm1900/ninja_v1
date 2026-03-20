package game

type Nature string

const (
	NatureFire      Nature = "fire"
	NatureWind      Nature = "wind"
	NatureLightning Nature = "lightning"
	NatureEarth     Nature = "earth"
	NatureWater     Nature = "water"
	NatureYin       Nature = "yin"
	NatureYang      Nature = "yang"
)

type NatureSet string

const (
	NsFire      NatureSet = "fire"
	NsWind      NatureSet = "wind"
	NsLightning NatureSet = "lightning"
	NsEarth     NatureSet = "earth"
	NsWater     NatureSet = "water"
	NsYin       NatureSet = "yin"
	NsYang      NatureSet = "yang"

	NsScorch    NatureSet = "scorch"
	NsLava      NatureSet = "lava"
	NsBoil      NatureSet = "boil"
	NsMagnet    NatureSet = "magnet"
	NsIce       NatureSet = "ice"
	NsExplosion NatureSet = "explosion"
	NsStorm     NatureSet = "storm"
	NsWood      NatureSet = "wood"
	NsYinYang   NatureSet = "yinyang"
	NsDust      NatureSet = "dust"
)

var NATURES = map[NatureSet][]Nature{
	NsFire:      {NatureFire},
	NsWind:      {NatureWind},
	NsLightning: {NatureLightning},
	NsEarth:     {NatureEarth},
	NsWater:     {NatureWater},
	NsYin:       {NatureYin},
	NsYang:      {NatureYang},

	NsScorch: {NatureFire, NatureWind},
	// ??? {NatureFire, NatureLightning}
	NsLava: {NatureFire, NatureEarth},
	NsBoil: {NatureFire, NatureWater},
	// ??? {NatureWind, NatureLightning}
	NsMagnet:    {NatureWind, NatureEarth},
	NsIce:       {NatureWind, NatureWater},
	NsExplosion: {NatureLightning, NatureEarth},
	NsStorm:     {NatureLightning, NatureWater},
	NsWood:      {NatureEarth, NatureWater},
	NsYinYang:   {NatureYin, NatureYang},
	NsDust:      {NatureFire, NatureEarth, NatureWind},
}
