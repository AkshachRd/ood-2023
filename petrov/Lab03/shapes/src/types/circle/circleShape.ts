import { ShapeVisitor } from '../shape/shapeVisitor';
import { ICircleShape } from './circle';

export class CircleShape implements ICircleShape {
    private radius: number;
    private x: number;
    private y: number;
    private type: string = 'circle';

    constructor(x: number, y: number, radius: number) {
        this.radius = radius;
        this.x = x;
        this.y = y;
    }

    public getRadius(): number {
        return this.radius;
    }

    public setRadius(radius: number): void {
        this.radius = radius;
    }

    public getX(): number {
        return this.x;
    }

    public setX(x: number): void {
        this.x = x;
    }

    public getY(): number {
        return this.y;
    }

    public setY(y: number): void {
        this.y = y;
    }

    public getType(): string {
        return this.type;
    }

    public draw(canvasContext: CanvasRenderingContext2D): void {
        canvasContext.beginPath();
        canvasContext.arc(this.x, this.y, this.radius, 0, 2 * Math.PI, false);
        canvasContext.fillStyle = 'green';
        canvasContext.fill();
        canvasContext.lineWidth = 5;
        canvasContext.strokeStyle = '#003300';
        canvasContext.stroke();
    }

    public accept(shapeVisitor: ShapeVisitor): void {
        shapeVisitor.visitCircleShape(this);
    }
}
