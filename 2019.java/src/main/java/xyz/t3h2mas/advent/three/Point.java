package xyz.t3h2mas.advent.three;

import java.util.Objects;

public class Point {
    private final int x;
    private final int y;


    public Point(int x, int y) {
        this.x = x;
        this.y = y;
    }

    public Point moveLeft() {
        return new Point(this.x - 1, this.y);
    }

    public Point moveRight() {
        return new Point(this.x + 1, this.y);
    }

    public Point moveUp() {
        return new Point(this.x, this.y + 1);
    }

    public Point moveDown() {
        return new Point(this.x, this.y - 1);
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Point point = (Point) o;
        return x == point.x &&
                y == point.y;
    }

    @Override
    public int hashCode() {
        return Objects.hash(x, y);
    }

    public int getX() {
        return x;
    }

    public int getY() {
        return y;
    }

    @Override
    public String toString() {
        return "Point{" +
                "x=" + x +
                ", y=" + y +
                '}';
    }
}
