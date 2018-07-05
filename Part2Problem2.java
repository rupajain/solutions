package problemsolving;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.ArrayList;
import java.util.Arrays;
import java.util.Scanner;

public class Part2Problem2 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc = new Scanner(System.in);
		System.out.println("please enter file location...");
		String sfileloc = sc.nextLine();
		Part2Problem2 pobj = new Part2Problem2();
		FileReader freader;
		int finalsum = 0;
		int div = 0;
		try {
			freader = new FileReader(sfileloc);

			BufferedReader bf = new BufferedReader(freader);
			String row = bf.readLine();

			String[] s1 = null;

			while (row != null) {
				s1 = row.split(",");
				div = pobj.divResult(s1);
				System.out.println("rows is..." + "...." + s1[0]);
				row = bf.readLine();

				finalsum = finalsum + div;
			}

			System.out.println("final sum is...." + finalsum);
			bf.close();
			sc.close();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (NumberFormatException e) {
			// TODO: handle exception
			System.out.println(e.getMessage());
			e.printStackTrace();
		}

	}

	private int divResult(String[] row) {
		// TODO Auto-generated method stub

		int[] arr1 = new int[row.length];
		ArrayList<Integer> div1 = new ArrayList<>();
		int finaldev = 0;

		for (int i = 0; i < row.length; i++) {

			arr1[i] = Integer.parseInt((row[i]));
		}

		Arrays.sort(arr1);
		System.out.println("after sorting .." + Arrays.toString(arr1));

		for (int i = 0; i < arr1.length; i++) {
			for (int j = 1; j < arr1.length; j++) {
				if ((arr1[j] % arr1[i]) == 0 && (arr1[j] != arr1[i])) {
					finaldev = arr1[j] / arr1[i];
					div1.add(finaldev);
					System.out.println("equal div..." + arr1[j] + "   " + arr1[i] + "final res is" + finaldev);
				}
			}

		}

		return finaldev;

	}

}
