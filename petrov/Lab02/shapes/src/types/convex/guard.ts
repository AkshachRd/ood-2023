import { Shape } from "../shape/shape"
import { ConvexShape } from "./convexShape"

export const isConvexShape = (shape: Shape): shape is ConvexShape => {
    return shape.getType() === "convex" &&
        (shape as ConvexShape).getPoints !== undefined &&
        (shape as ConvexShape).setPoints !== undefined
}