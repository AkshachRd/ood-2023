import { ICircleShape } from '../../circle/circle';
import { CircleShape } from '../../circle/circleShape';

export class CircleDecorator implements ICircleShape {
    protected circle: CircleShape;

    constructor(circle: CircleShape) {
        this.circle = circle;
    }

    public getRadius(): number {
        return this.circle.getRadius();
    }

    public setRadius(radius: number): void {
        this.circle.setRadius(radius);
    }

    public getX(): number {
        return this.circle.getX();
    }

    public setX(x: number): void {
        this.circle.setX(x);
    }

    public getY(): number {
        return this.circle.getY();
    }

    public setY(y: number): void {
        this.circle.setY(y);
    }

    public getType(): string {
        return this.circle.getType();
    }

    public getPerimeter(): number {
        return 2 * Math.PI * this.circle.getRadius();
    }

    public getArea(): number {
        return Math.PI * Math.pow(this.circle.getRadius(), 2);
    }

    draw(canvasContext: CanvasRenderingContext2D): void {
        this.circle.draw(canvasContext);
    }
}
