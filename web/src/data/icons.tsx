import type { ActorBaseStat } from '#/lib/game/actor'
import { FaPrayingHands } from 'react-icons/fa'
import { GiHearts, GiSprint, GiFist } from 'react-icons/gi'
import { MdEnergySavingsLeaf } from 'react-icons/md'

import type { IconType } from 'react-icons/lib'

const Ninjutsu: IconType = (props) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    stroke="currentColor"
    fill="none"
    strokeWidth="0"
    viewBox="0 0 64 64"
    height="1em"
    width="1em"
    {...props}
  >
    <title>Ninjutsu</title>
    <g
      stroke="currentColor"
      strokeWidth="3.5"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M22 8h19l7 8v22c0 9-6 15-16 18c-10-3-16-9-16-18V8h6" />
      <path d="M41 8v8h7" />
      <path d="M26 16h12" />
      <path d="M22 24c5-2 15-2 20 0" />
      <path d="M32 16v29" />
      <path d="M20 32c6-2 18-2 24 0" />
      <path d="M22 40c5 2 15 2 20 0" />
      <path d="M27 47c3 2 6 4 10 6" />
      <path d="M18 51c7-2 12-6 14-11" />
      <path d="M46 48c-5 1-9 4-12 8" />
    </g>
  </svg>
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

const Taijutsu: IconType = (props) => (
  <svg
    xmlns="http://www.w3.org/2000/svg"
    stroke="currentColor"
    fill="none"
    strokeWidth="0"
    viewBox="0 0 64 64"
    height="1em"
    width="1em"
    {...props}
  >
    <title>Taijutsu</title>
    <g
      stroke="currentColor"
      strokeWidth="4"
      strokeLinecap="round"
      strokeLinejoin="round"
    >
      <path d="M22 29V18a3 3 0 0 1 6 0v9" />
      <path d="M28 27V14a3 3 0 0 1 6 0v13" />
      <path d="M34 27V16a3 3 0 0 1 6 0v11" />
      <path d="M40 29v-5a3 3 0 0 1 6 0v12c0 11-8 19-19 19h-3c-8 0-14-6-14-14v-9a3 3 0 0 1 6 0v4" />
      <path d="M22 39c4 2 9 3 14 3c4 0 7 0 10-1" />
      <path d="M15 49l-6 2" />
      <path d="M49 49l6 2" />
    </g>
  </svg>
)

const STAT_ICONS: Record<ActorBaseStat, IconType | null> = {
  accuracy: null,
  evasion: null,
  genjutsu: Genjutsu,
  hp: GiHearts,
  ninjutsu: Ninjutsu, // FaPrayingHands,
  speed: GiSprint,
  stamina: MdEnergySavingsLeaf,
  taijutsu: GiFist,
}

export { STAT_ICONS }
