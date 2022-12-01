package xyz.t3h2mas.advent.one.domain;

public class FuelForModules implements FuelStrategy {
    private int calculateRequiredFuelFor(int mass) {
        return (int) Math.floor((float) mass / 3.0) - 2;
    }

    @Override
    public int fuelRequiredBy(SpaceCraft sc) {
        return sc.getModules()
                .stream()
                .map(m -> this.calculateRequiredFuelFor(m.getMass()))
                .reduce(0, Integer::sum);
    }
}
