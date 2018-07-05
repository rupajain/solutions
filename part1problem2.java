package problemsolving;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class part1problem2 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		try {

			ArrayList<Integer> maxmindiff = new ArrayList<Integer>();
			int[] rmaxmin = new int[2];
			part1problem2 pobj = new part1problem2();
			Scanner sc = new Scanner(System.in);
			System.out.println("please enter file location...");
			String sfileloc = sc.nextLine();

			FileReader freader = new FileReader(sfileloc);
			BufferedReader bf = new BufferedReader(freader);
			String row = bf.readLine();

			System.out.println("rows.." + row);

			int noofrows = 0;
			String[] s1 = null;
			while (row != null) {
				s1 = row.split("/t");

				System.out.println("number of rows are..." + (noofrows + 1) + "...." + s1[0]);

				rmaxmin = pobj.findMaxMin(s1[0]);
				maxmindiff.add(rmaxmin[0] - rmaxmin[1]);

				row = bf.readLine();
				noofrows++;
			}

			int sum = 0;
			for (int i = 0; i < maxmindiff.size(); i++) {

				sum = sum + maxmindiff.get(i);
			}
			System.out.println("sum is..." + sum);
			bf.close();
			sc.close();

		} catch (NumberFormatException e) {
			// TODO: handle exception
			System.out.println("please enter...." + e.getMessage() + ".....");
			e.printStackTrace();
		} catch (FileNotFoundException e) {
			// TODO: handle exception
			System.out.println("File not found...." + e.getMessage());
		} catch (IOException e) {
			// TODO: handle exception
			System.out.println(e.getMessage());
		}

	}

	public int[] findMaxMin(String row) {

		String[] s1 = row.split("\t");
		int[] maxmin = new int[2];

		int[] arr1 = new int[s1.length];
		System.out.println("array is..." + Arrays.toString(s1) + "   " + "length is.." + s1.length);

		for (int i = 0; i < s1.length; i++) {

			arr1[i] = Integer.parseInt((s1[i]));
		}

		System.out.println("array is..." + Arrays.toString(arr1));
		int max = 0;
		int min = 0;
		int tempmax = 0;
		int tempmin = 0;
		tempmax = arr1[0];
		tempmin = arr1[0];
		for (int i = 1; i < arr1.length; i++) {

			if (tempmax < arr1[i]) {
				tempmax = arr1[i];

			}
			max = tempmax;
			if (tempmin > arr1[i]) {
				tempmin = arr1[i];
			}
			min = tempmin;

		}

		System.out.println("max of row.." + max + "    min of row..." + tempmin);

		maxmin[0] = max;
		maxmin[1] = min;
		return maxmin;
	}

}
