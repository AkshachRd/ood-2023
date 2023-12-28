export abstract class Shape {
    abstract getType(): string;
    abstract draw(canvasContext: CanvasRenderingContext2D): void;
}