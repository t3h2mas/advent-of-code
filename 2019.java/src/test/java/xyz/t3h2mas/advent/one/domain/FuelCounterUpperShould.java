package xyz.t3h2mas.advent.one.domain;

import org.junit.Test;
import org.junit.runner.RunWith;
import org.mockito.Mock;
import org.mockito.junit.MockitoJUnitRunner;

import static org.mockito.Mockito.verify;

@RunWith(MockitoJUnitRunner.class)
public class FuelCounterUpperShould {
    @Mock
    FuelStrategy fuelStrategy;

    @Test
    public void use_a_fuel_strategy_to_calculate_required_fuel() {
        FuelCounterUpper sut = new FuelCounterUpper(fuelStrategy);
        SpaceCraft sc = new SpaceCraft();

        sut.calculateFuelRequirements(sc);

        verify(fuelStrategy).fuelRequiredBy(sc);
    }
}