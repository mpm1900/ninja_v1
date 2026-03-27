import type { ActorBaseStat } from '#/lib/game/actor'
import { GiHearts, GiSprint, GiPunch, GiMagicSwirl } from 'react-icons/gi'
import { MdEnergySavingsLeaf } from 'react-icons/md'

import type { IconType } from 'react-icons/lib'

const Akatsuki: IconType = (props) => (
  <img src="/icons/akatsuki.svg" alt="Akatsuki" {...(props as any)} />
)
const Hatake: IconType = (props) => (
  <img src="/icons/hatake.svg" alt="Hatake" {...(props as any)} />
)
const Konoha: IconType = (props) => (
  <img src="/icons/konoha.svg" alt="Konoha" {...(props as any)} />
)
const Uchiha: IconType = (props) => (
  <img src="/icons/uchiha.svg" alt="Uchiha" {...(props as any)} />
)
const Uzumaki: IconType = (props) => (
  <img src="/icons/uzumaki.svg" alt="Uzumaki" {...(props as any)} />
)

const Genjutsu: IconType = (props) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    stroke="currentColor"
    fill="currentColor"
    strokeWidth="0"
    viewBox="0 0 64 64"
    height="1em"
    width="1em"
    {...props}
  >
    <title>Genjutsu</title>
    <g fill="currentColor">
      <path
        fillRule="evenodd"
        clipRule="evenodd"
        d="M3 32s12-19 29-19 29 19 29 19-12 19-29 19S3 32 3 32Zm44 0a15 15 0 1 1-30 0a15 15 0 0 1 30 0Zm-15-9c-5 0-10 3-15 9c5 6 10 9 15 9s10-3 15-9c-5-6-10-9-15-9Zm0-4c7 1 13 6 13 13s-6 12-13 13"
      />
    </g>
  </svg>
)

const STAT_ICONS: Record<ActorBaseStat, IconType | undefined> = {
  accuracy: undefined,
  evasion: undefined,
  genjutsu: Genjutsu,
  hp: GiHearts,
  ninjutsu: GiMagicSwirl,
  speed: GiSprint,
  chakra: MdEnergySavingsLeaf,
  taijutsu: GiPunch,
}

const SHINOBI_ICONS: Record<string, IconType | undefined> = {
  akatsuki: Akatsuki,
  hatake: Hatake,
  konoha: Konoha,
  uchiha: Uchiha,
  uzumaki: Uzumaki,
}

export { Akatsuki, STAT_ICONS, SHINOBI_ICONS }
