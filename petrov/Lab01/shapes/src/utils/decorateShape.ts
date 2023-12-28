import { CircleShape } from "../types/circle/circleShape";
import { ConvexShape } from "../types/convex/convexShape";
import { ConcreteCircleDecorator } from "../types/decorators/circleDecorator/concreteCircleDecorator";
import { ConcreteConvexDecorator } from "../types/decorators/convexDecorator/concreteConvexDecorator";
import { ConcreteRectangleDecorator } from "../types/decorators/rectangleDecorator/concreteRectangleDecorator";
import { RectangleShape } from "../types/rectangle/rectangleShape";
import { Shape } from "../types/shape/shape";

export const decorateShape = (shape: Shape) => {
    if (shape.getType() === 'circle') {
      return new ConcreteCircleDecorator(shape as CircleShape);
    } else if (shape.getType() === 'rectangle') {
      return new ConcreteRectangleDecorator(shape as RectangleShape);
    } else if (shape.getType() === 'convex') {
      return new ConcreteConvexDecorator(shape as ConvexShape);
    }
  };