import { Point } from "../point/point";
import { Shape } from "../shape/shape";

export interface IConvexShape extends Shape {
    getPoints(): Point[];
    setPoints(points: Point[]): void;
    getType(): string;
    draw(canvasContext: CanvasRenderingContext2D): void;
}