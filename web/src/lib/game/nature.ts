type Nature =
  | 'pure'
  | 'fire'
  | 'wind'
  | 'lightning'
  | 'earth'
  | 'water'
  | 'yang'
  | 'yin'
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
  | 'jashin'

const natureNames: Partial<Record<NatureSet, string>> = {
  pure: '纯',
  fire: '火',
  wind: '風',
  lightning: '雷',
  earth: '土',
  water: '水',
  yin: '陰',
  yang: '陽',
  wood: '木',
  yinyang: '陰陽',
  jashin: '邪',
}

const natureIndexes: Record<NatureSet, number> = {
  fire: 0,
  wind: 1,
  lightning: 2,
  earth: 3,
  water: 4,
  yin: 5,
  yang: 6,
  pure: 7,
  scorch: 8,
  lava: 9,
  boil: 10,
  magnet: 12,
  ice: 13,
  explosion: 14,
  storm: 15,
  wood: 16,
  yinyang: 17,
  dust: 18,
  jashin: 19,
}

export type { Nature, NatureSet }
export { natureNames, natureIndexes }
