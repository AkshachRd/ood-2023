import { IRectangleShape } from "./rectangle";

export class RectangleShape implements IRectangleShape {
    private width: number;
    private height: number;
    private y: number;
    private x: number;
    private type: string = 'rectangle';

    constructor(x: number, y: number, width: number, height: number) {
        this.width = width;
        this.height = height;
        this.y = y;
        this.x = x;
    }

    public getWidth(): number {
        return this.width;
    }

    public setWidth(width: number): void {
        this.width = width;
    }

    public getHeight(): number {
        return this.height;
    }

    public setHeight(height: number): void {
        this.height = height;
    }

    public getY(): number {
        return this.y;
    }

    public setY(y: number): void {
        this.y = y;
    }

    public getX(): number {
        return this.x;
    }

    public setX(x: number): void {
        this.x = x;
    }

    public getType(): string {
        return this.type;
    }

    public draw(canvasContext: CanvasRenderingContext2D): void {
        canvasContext.fillRect(this.x, this.y, this.width, this.height);
    }
}
