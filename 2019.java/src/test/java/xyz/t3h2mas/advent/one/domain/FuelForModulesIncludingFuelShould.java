package xyz.t3h2mas.advent.one.domain;

import org.junit.Test;

import static org.junit.Assert.assertEquals;

public class FuelForModulesIncludingFuelShould {
    private final int INPUT_MASS = 1969;
    private final int EXPECTED_REQUIRED_FUEL = 966;

    @Test
    public void calculate_fuel_required_for_modules_and_fuel_of_a_spacecraft() {
        FuelForModulesIncludingFuel sut = new FuelForModulesIncludingFuel();

        SpaceCraft spaceCraft = new SpaceCraft();
        spaceCraft.addModule(new Module(INPUT_MASS));

        assertEquals(sut.fuelRequiredBy(spaceCraft), EXPECTED_REQUIRED_FUEL);
    }
}
