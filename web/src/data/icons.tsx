import { cn } from '#/lib/utils'
import {
  GiComa,
  GiFlamer,
  GiLightningTrio,
  GiShieldcomb,
  GiStarSwirl,
  GiNoodles,
  GiKevlarVest,
  GiDoubled,
  GiBeastEye,
  GiHealing,
  GiTopaz,
  GiMinotaur,
  GiNightSleep,
  GiHandheldFan,
  GiSandstorm,
  GiHealthIncrease,
  GiWaterRecycling,
  GiPoisonBottle,
} from 'react-icons/gi'
import {
  PiCaretDoubleUpDuotone,
  PiOnigiriFill,
  PiArrowFatLinesDownBold,
} from 'react-icons/pi'
import type { IconType } from 'react-icons/lib'
import { MdFileUploadOff } from 'react-icons/md'
import { GrFastForward } from 'react-icons/gr'
import { TbTagPlus } from 'react-icons/tb'
import { HiScale } from 'react-icons/hi2'
import { FaFrog } from 'react-icons/fa6'
import { TbScanEye } from 'react-icons/tb'

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
  coerced: GiComa,
  cmo_chakra: ({ className, ...props }) => (
    <GiDoubled className={cn('text-indigo-400', className)} {...props} />
  ),
  cmo_speed: ({ className, ...props }) => (
    <GiDoubled className={cn('text-emerald-300', className)} {...props} />
  ),
  cmo_strength: ({ className, ...props }) => (
    <GiDoubled className={cn('text-orange-300', className)} {...props} />
  ),
  std_chakra: ({ className, ...props }) => (
    <PiArrowFatLinesDownBold
      className={cn('text-indigo-400', className)}
      {...props}
    />
  ),
  std_speed: ({ className, ...props }) => (
    <PiArrowFatLinesDownBold
      className={cn('text-emerald-300', className)}
      {...props}
    />
  ),
  std_strength: ({ className, ...props }) => (
    <PiArrowFatLinesDownBold
      className={cn('text-orange-300', className)}
      {...props}
    />
  ),
  fast_thinking: GrFastForward,
  gedo_shard: GiTopaz,
  guts: GiMinotaur,
  haze: HiScale,
  healing_tactics: GiHealing,
  ichiraku_ramen: GiNoodles,
  inner_focus: TbScanEye,
  intimidate: GiBeastEye,
  onigiri: PiOnigiriFill,
  paralyzed: ({ className, ...props }) => (
    <GiLightningTrio className={cn('text-yellow-400', className)} {...props} />
  ),
  poisoned: ({ className, ...props }) => (
    <GiPoisonBottle className={cn('text-lime-500', className)} {...props} />
  ),
  priority_failure: MdFileUploadOff,
  protected: GiShieldcomb,
  regeneration: GiHealthIncrease,
  sage_mode: FaFrog,
  sand_aura: GiSandstorm,
  seal_up: TbTagPlus,
  shinobi_vest: GiKevlarVest,
  sleeping: GiNightSleep,
  speed_up: ({ className, ...props }) => (
    <PiCaretDoubleUpDuotone
      className={cn('text-emerald-400', className)}
      {...props}
    />
  ),
  stunned: GiStarSwirl,
  taunted: Anger,
  uchiha_fan: GiHandheldFan,
  water_absorb: GiWaterRecycling,
}

export { Akatsuki, SHINOBI_ICONS, MODIFIER_ICONS }
