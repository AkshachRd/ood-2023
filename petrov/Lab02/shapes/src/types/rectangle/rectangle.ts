import { Shape } from "../shape/shape";

export interface IRectangleShape extends Shape {
    getWidth(): number;
    setWidth(width: number): void;
    getHeight(): number;
    setHeight(height: number): void;
    getY(): number;
    setY(y: number): void;
    getX(): number;
    setX(x: number): void;
    getType(): string;
    draw(canvasContext: CanvasRenderingContext2D): void;
}