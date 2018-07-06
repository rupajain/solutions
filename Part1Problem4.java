package problemsolving;

import java.io.BufferedReader;
import java.io.FileNotFoundException;
import java.io.FileReader;
import java.io.IOException;
import java.util.HashSet;
import java.util.Scanner;
import java.util.Set;

public class Part1Problem4 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		System.out.println("please enter file location..");
		Scanner sc = new Scanner(System.in);
		String fileloc = sc.nextLine();
		String s = null;
		int count = 0;
		FileReader freader=null;
		BufferedReader bf=null;
		try {
			 freader = new FileReader(fileloc);
			 bf = new BufferedReader(freader);

			s = bf.readLine();
		
		while (s != null) {

			String[] s1 = s.split(" ");
			Set<String> words = new HashSet<String>();
			for (int i = 0; i < s1.length; i++) {
				System.out.println("words from string are..." + s1[i]);
				words.add(s1[i]);

			}
			System.out.println("words from set of words are..." + words);

			System.out.println("diff.." + words.size() + "....." + s1.length);
			if (s1.length != words.size()) {
				System.out.println("is not valid");
			} else {
				count++;
				System.out.println("is valid" + count);
			}
			s = bf.readLine();
		}
		} catch (FileNotFoundException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		} catch (IOException e) {
			// TODO Auto-generated catch block
			e.printStackTrace();
		}
	}

}
