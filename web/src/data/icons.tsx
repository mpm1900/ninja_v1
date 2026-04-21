import { cn } from '#/lib/utils'
import {
  GiFlamer,
  GiLightningTrio,
  GiShieldcomb,
  GiStarSwirl,
  GiNoodles,
  GiKevlarVest,
  GiDoubled,
  GiBeastEye,
  GiHealing,
} from 'react-icons/gi'
import { PiCaretDoubleUpDuotone } from "react-icons/pi";
import type { IconType } from 'react-icons/lib'
import { WiSandstorm } from "react-icons/wi";

const Aburame: IconType = (props) => (
  <img src="/icons/aburame.svg" alt="aburame" {...(props as any)} />
)
const Anger: IconType = (props) => (
  <img src="/icons/anger.svg" alt="Ame" {...(props as any)} />
)
const Ame: IconType = (props) => (
  <img src="/icons/ame.svg" alt="Ame" {...(props as any)} />
)
const Akatsuki: IconType = (props) => (
  <img src="/icons/akatsuki.svg" alt="Akatsuki" {...(props as any)} />
)
const Hatake: IconType = (props) => (
  <img src="/icons/hatake.svg" alt="Hatake" {...(props as any)} />
)
const Iwa: IconType = (props) => (
  <img src="/icons/iwa.svg" alt="Iwa" {...(props as any)} />
)
const Konoha: IconType = (props) => (
  <img src="/icons/konoha.svg" alt="Konoha" {...(props as any)} />
)
const Kumo: IconType = (props) => (
  <img src="/icons/kumo.svg" alt="Kumo" {...(props as any)} />
)
const Kuri: IconType = (props) => (
  <img src="/icons/kuri.svg" alt="Kuri" {...(props as any)} />
)
const Nara: IconType = (props) => (
  <img src="/icons/nara.svg" alt="Nara" {...(props as any)} />
)
const Oto: IconType = (props) => (
  <img src="/icons/oto.svg" alt="Oto" {...(props as any)} />
)
const Senju: IconType = (props) => (
  <img src="/icons/senju.svg" alt="Senju" {...(props as any)} />
)
const Sun: IconType = (props) => (
  <img src="/icons/sun.svg" alt="Sun" {...(props as any)} />
)
const Taki: IconType = (props) => (
  <img src="/icons/taki.svg" alt="Taki" {...(props as any)} />
)
const Uchiha: IconType = (props) => (
  <img src="/icons/uchiha.svg" alt="Uchiha" {...(props as any)} />
)
const Uzumaki: IconType = (props) => (
  <img src="/icons/uzumaki.svg" alt="Uzumaki" {...(props as any)} />
)
const Yuga: IconType = (props) => (
  <img src="/icons/yuga.svg" alt="Yuga" {...(props as any)} />
)

const SHINOBI_ICONS: Record<string, IconType> = {
  aburame: Aburame,
  ame: Ame,
  akatsuki: Akatsuki,
  hatake: Hatake,
  iwa: Iwa,
  konoha: Konoha,
  kumo: Kumo,
  kuri: Kuri,
  nara: Nara,
  oto: Oto,
  senju: Senju,
  sun: Sun,
  taki: Taki,
  uchiha: Uchiha,
  uzumaki: Uzumaki,
  yuga: Yuga,
}

const MODIFIER_ICONS: Record<string, IconType> = {
  burned: ({ className, ...props }) => (
    <GiFlamer className={cn('text-orange-400', className)} {...props} />
  ),
  paralyzed: ({ className, ...props }) => (
    <GiLightningTrio className={cn('text-yellow-400', className)} {...props} />
  ),
  protected: GiShieldcomb,
  stunned: GiStarSwirl,
  taunted: Anger,
  ichiraku_ramen: GiNoodles,
  shinobi_vest: GiKevlarVest,
  cmo_strength: ({ className, ...props }) => (
    <GiDoubled className={cn('text-orange-300', className)} {...props} />
  ),
  cmo_chakra: ({ className, ...props }) => (
    <GiDoubled className={cn('text-indigo-400', className)} {...props} />
  ),
  cmo_speed: ({ className, ...props }) => (
    <GiDoubled className={cn('text-emerald-300', className)} {...props} />
  ),
  speed_up: ({ className, ...props }) => (
    <PiCaretDoubleUpDuotone className={cn('text-emerald-400', className)} {...props} />
  ),
  sand_aura: WiSandstorm,
  intimidate: GiBeastEye,
  healing_tactics: GiHealing,
}

export { Akatsuki, SHINOBI_ICONS, MODIFIER_ICONS }
