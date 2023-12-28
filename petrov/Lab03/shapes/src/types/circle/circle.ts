import { Shape } from "../shape/shape";

export interface ICircleShape extends Shape {
    getRadius(): number;
    setRadius(radius: number): void;
    getX(): number;
    setX(x: number): void;
    getY(): number;
    setY(y: number): void;
    getType(): string;
    draw(canvasContext: CanvasRenderingContext2D): void;
}