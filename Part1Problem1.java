package problemsolving;

import java.util.Arrays;
import java.util.Scanner;

public class Part1Problem1 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		
		Scanner scr=new Scanner(System.in);
		// System.out.println("enter the length of array");
		 //System.out.println("Enter the number");
		String number=scr.nextLine();
		
		
		
		
		 int arrayone[]=new int[number.length()];
		 int newArray[]=new int[number.length()+1];
		 for(int i=0;i<number.length();i++)
		 {
		 arrayone[i]=Integer.valueOf(number.charAt(i))-'0';
		 }
		 System.out.println(Arrays.toString(arrayone));

		 //int count=length;
		 if(arrayone[0]==arrayone[arrayone.length-1])
		 {
		 int temp=arrayone[0];
		 newArray[newArray.length-1]=temp;
		 }


		 for(int j=1;j<number.length();j++)
		 {

		 if(arrayone[j-1]==arrayone[j] )
		 {
		 newArray[j]=arrayone[j];
		 }


		 }




		 System.out.println(Arrays.toString(newArray));

		 int sum=0;
		 for(int k=0;k<newArray.length;k++)
		 {

		 sum=sum+newArray[k];
		 }
		 System.out.println("sum of repeated letters are  "+sum);

	}

}
