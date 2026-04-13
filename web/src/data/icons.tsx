import { GiFlamer, GiLightningTrio, GiShieldcomb, GiStarSwirl } from 'react-icons/gi'
import type { IconType } from 'react-icons/lib'

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
  burned: GiFlamer,
  paralyzed: GiLightningTrio,
  protected: GiShieldcomb,
  stunned: GiStarSwirl,
  taunted: Anger,
}
const MODIFIER_CLASSES: Record<string, string> = {
  burned: 'text-orange-400',
  paralyzed: 'text-yellow-400',
  protected: '',
}

export { Akatsuki, SHINOBI_ICONS, MODIFIER_ICONS, MODIFIER_CLASSES }
