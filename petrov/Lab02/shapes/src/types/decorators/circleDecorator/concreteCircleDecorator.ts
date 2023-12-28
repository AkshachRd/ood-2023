import { CircleShape } from '../../circle/circleShape';
import { IConcreteDecorator } from '../concreteDecorator';
import { CircleDecorator } from './circleDecorator';

export class ConcreteCircleDecorator extends CircleDecorator implements IConcreteDecorator {
    constructor(circle: CircleShape) {
        super(circle)
    }

    public getPerimeter(): number {
        return 2 * Math.PI * this.circle.getRadius();
    }

    public getArea(): number {
        return Math.PI * Math.pow(this.circle.getRadius(), 2);
    }
}
