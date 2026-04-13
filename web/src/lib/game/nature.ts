type Nature =
  | 'tai'
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
  | 'particle'
  | 'jashin'

const natureWeakness: Record<Nature, Nature | undefined> = {
  tai: undefined,
  pure: undefined,
  fire: 'water',
  wind: 'fire',
  lightning: 'wind',
  earth: 'lightning',
  water: 'earth',
  yang: undefined,
  yin: undefined,
}
const natureResistance: Record<Nature, Nature | undefined> = {
  tai: undefined,
  pure: undefined,
  fire: 'wind',
  wind: 'lightning',
  lightning: 'earth',
  earth: 'water',
  water: 'fire',
  yang: undefined,
  yin: undefined,
}
const natureSetMap: Record<NatureSet, Array<Nature>> = {
  tai: ['tai'],
  pure: ['pure'],
  fire: ['fire'],
  wind: ['wind'],
  lightning: ['lightning'],
  earth: ['earth'],
  water: ['water'],
  yang: ['yang'],
  yin: ['yin'],
  scorch: ['fire', 'wind'],
  lava: ['fire', 'earth'],
  boil: ['fire', 'water'],
  magnet: ['wind', 'earth'],
  ice: ['wind', 'water'],
  explosion: ['earth', 'lightning'],
  storm: ['lightning', 'water'],
  wood: ['earth', 'water'],
  yinyang: ['yin', 'yang'],
  particle: ['fire', 'earth', 'lightning'],
  jashin: [],
}

const natureNames: Partial<Record<NatureSet, string>> = {
  tai: '体',
  pure: '纯',
  fire: '火',
  wind: '風',
  lightning: '雷',
  earth: '土',
  water: '水',
  yin: '陰',
  yang: '陽',
  ice: '氷',
  explosion: '爆',
  storm: '嵐',
  wood: '木',
  yinyang: '陰陽',
  particle: '塵',
  jashin: '邪',
}

const natureIndexes: Record<NatureSet, number> = {
  tai: -1,
  fire: 0,
  wind: 1,
  lightning: 2,
  earth: 3,
  water: 4,
  yin: 5,
  yang: 6,
  scorch: 8,
  lava: 9,
  boil: 10,
  magnet: 12,
  ice: 13,
  explosion: 14,
  storm: 15,
  wood: 16,
  yinyang: 17,
  particle: 18,
  pure: 19,
  jashin: 20,
}

function getWeakness(...natures: Array<NatureSet>): Array<Nature> {
  const list = natures.flatMap((nature) => {
    const base = natureSetMap[nature]
    return base.map((n) => natureWeakness[n]).filter((n) => n !== undefined)
  })
  return Array.from(new Set(list))
}

function getResistance(...natures: Array<NatureSet>): Array<Nature> {
  const list = natures.flatMap((nature) => {
    const base = natureSetMap[nature]
    return base.map((n) => natureResistance[n]).filter((n) => n !== undefined)
  })
  return Array.from(new Set(list))
}

export type { Nature, NatureSet }
export {
  natureNames,
  natureIndexes,
  natureSetMap,
  natureWeakness,
  natureResistance,
  getWeakness,
  getResistance,
}
