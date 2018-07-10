package com.exilant.day2.stats;

import java.util.Comparator;
import java.util.List;
import java.util.Map;
import java.util.stream.Collectors;
import java.util.stream.Stream;

public class DragonsClient {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		List<Dragon> dragonList = Stream.of(new Dragon("Errol", "USA", 15), new Dragon("Errol", "USA", 56),
				new Dragon("Firnan", "USA", 5), new Dragon("Firnan", "France", 15), new Dragon("IceFire", "France", 15),
				new Dragon("IceFire", "USA", 15)).collect(Collectors.toList());

		// to find number of dragons
		Map<String, Long> mapDragobnCount = dragonList.stream()
				.collect(Collectors.groupingBy(Dragon::getName, Collectors.counting()));

		System.out.println(mapDragobnCount);
		Map<String, Integer> sumofDragons = dragonList.stream()
				.collect(Collectors.groupingBy(Dragon::getName, Collectors.summingInt(Dragon::getCount)));

		System.out.println(".............report of dragons");
		sumofDragons.forEach((k, v) -> {
			System.out.println("Dragon Name " + k + " number of dragons " + v);
		});
//
		
		System.out.println("sum of dragons....."+dragonList.stream().collect(Collectors.summingInt(Dragon::getCount)));
		System.out.println("sum of dragons....."+dragonList.stream().collect(Collectors.averagingInt((Dragon::getCount))));
		//System.out.println("sum of dragons....."+dragonList.stream().collect(Collectors.minBy(Comparator<Integer> c)));
		
		
	}
}
