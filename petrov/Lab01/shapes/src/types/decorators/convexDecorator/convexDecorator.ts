import { Point } from '../../point/point';
import { ConvexShape } from '../../convex/convexShape';
import { IConvexShape } from '../../convex/convex';

export class ConvexDecorator implements IConvexShape {
    protected convex: ConvexShape;

    constructor(convex: ConvexShape) {
        this.convex = convex;
    }

    public getPoints(): Point[] {
        return this.convex.getPoints();
    }

    public setPoints(points: Point[]): void {
        this.convex.setPoints(points);
    }

    public getType(): string {
        return this.convex.getType();
    }

    public getPerimeter(): number {
        const points = this.convex.getPoints();
        let perimeter = 0;

        for (let i = 0; i < points.length - 1; i++) {
            const point1 = points[i];
            const point2 = points[i + 1];
            perimeter += Math.sqrt(Math.pow(point2.x - point1.x, 2) + Math.pow(point2.y - point1.y, 2));
        }

        return perimeter;
    }

    public getArea(): number {
        const points = this.convex.getPoints();
        let area = 0;

        for (let i = 0; i < points.length - 1; i++) {
            const point1 = points[i];
            const point2 = points[i + 1];
            area += (point1.x * point2.y) - (point2.x * point1.y);
        }

        return Math.abs(area) / 2;
    }

    draw(canvasContext: CanvasRenderingContext2D): void {
        this.convex.draw(canvasContext);
    }
}
