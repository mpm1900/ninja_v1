import React, { useEffect, useRef } from 'react'

declare global {
  interface Window {
    VANTA: any
  }
}

export const VantaBackground = () => {
  const vantaRef = useRef<HTMLDivElement>(null)
  const vantaEffectRef = useRef<any>(null)

  useEffect(() => {
    const initVanta = () => {
      if (window.VANTA && window.VANTA.FOG && vantaRef.current && !vantaEffectRef.current) {
        try {
          vantaEffectRef.current = window.VANTA.FOG({
            el: vantaRef.current,
            mouseControls: true,
            touchControls: true,
            gyroControls: false,
            minHeight: 200.00,
            minWidth: 200.00,
            highlightColor: 0x8089,
            midtoneColor: 0xffa600,
            lowlightColor: 0xffffff,
            baseColor: 0x0,
            blurFactor: 0.35,
            speed: 0.3,
            zoom: 0.40
          })
        } catch (err) {
          console.error('Vanta initialization failed:', err)
        }
      }
    }

    // Check every 100ms if VANTA is available, up to 50 times (5 seconds)
    let attempts = 0
    const interval = setInterval(() => {
      if (window.VANTA && window.VANTA.FOG) {
        initVanta()
        clearInterval(interval)
      }
      attempts++
      if (attempts > 50) {
        clearInterval(interval)
      }
    }, 100)

    return () => {
      clearInterval(interval)
      if (vantaEffectRef.current) {
        vantaEffectRef.current.destroy()
        vantaEffectRef.current = null
      }
    }
  }, [])

  return (
    <div
      ref={vantaRef}
      style={{
        position: 'fixed',
        top: 0,
        left: 0,
        right: 0,
        bottom: 0,
        zIndex: -1,
        pointerEvents: 'none',
      }}
    />
  )
}
