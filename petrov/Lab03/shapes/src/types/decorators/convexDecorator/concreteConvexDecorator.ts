import { ConvexShape } from '../../convex/convexShape';
import { IConcreteDecorator } from '../concreteDecorator';
import { ConvexDecorator } from './convexDecorator';

export class ConcreteConvexDecorator extends ConvexDecorator implements IConcreteDecorator {
    constructor(convex: ConvexShape) {
        super(convex)
    }

    public getPerimeter(): number {
        const points = this.convex.getPoints();
        let perimeter = 0;

        for (let i = 0; i < points.length; i++) {
            const point1 = points[i];
            const point2 = points[(i + 1) % points.length]; // Wrap around to the first point for the last side
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
}
