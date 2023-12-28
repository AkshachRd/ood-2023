import { useEffect, useRef } from 'react'
import { Shape } from '../../types/shape/shape';

type CanvasProps = {
  shapes: Shape[];
}

export const Canvas = ({shapes}: CanvasProps) => {
  const canvasRef = useRef<HTMLCanvasElement>(null)

  useEffect(() => {
    const canvas = canvasRef.current
    if (canvas !== null) {
      const context = canvas.getContext('2d') as CanvasRenderingContext2D
      //Our first draw
      context.fillStyle = '#000000'
    }
  }, [])

  useEffect(() => {
    const canvas = canvasRef.current
    if (canvas !== null) {
      const context = canvas.getContext('2d') as CanvasRenderingContext2D
      context.clearRect(0, 0, canvas.width, canvas.height)
      shapes.forEach(shape => {
        shape.draw(context)
      })
    }
  }, [shapes])

  return <canvas width={1000} height={1000} ref={canvasRef} />
}