import { CircleShape } from "../circle/circleShape";
import { ConvexShape } from "../convex/convexShape";
import { RectangleShape } from "../rectangle/rectangleShape";

export interface ShapeVisitor {
    visitConvexShape(convexShape: ConvexShape): void;
    visitRectangleShape(rectangleShape: RectangleShape): void;
    visitCircleShape(circleShape: CircleShape): void;
}