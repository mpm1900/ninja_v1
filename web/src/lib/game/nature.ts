type Nature = 'fire' | 'wind' | 'lightning' | 'earth' | 'water' | 'yang' | 'yin'
type NatureSet =
  | Nature
  | 'scorch'
  | 'lava'
  | 'boil'
  | 'magnet'
  | 'ice'
  | 'explosion'
  | 'storm'
  | 'wood'
  | 'yinyang'
  | 'dust'

const natureNames: Partial<Record<NatureSet, string>> = {
  fire: '火',
  wind: '風',
  lightning: '雷',
  earth: '土',
  water: '水',
  yin: '陰',
  yang: '陽',
  wood: '木',
  yinyang: '陰陽',
}

const natureIndexes: Record<NatureSet, number> = {
  fire: 0,
  wind: 1,
  lightning: 2,
  earth: 3,
  water: 4,
  yin: 5,
  yang: 6,
  scorch: 7,
  lava: 8,
  boil: 9,
  magnet: 12,
  ice: 13,
  explosion: 14,
  storm: 15,
  wood: 16,
  yinyang: 17,
  dust: 18,
}

export type { Nature, NatureSet }
export { natureNames, natureIndexes }
