package problemsolving;

import java.util.Arrays;
import java.util.Scanner;

public class Part2Problem1 {

	public static void main(String[] args) {
		// TODO Auto-generated method stub
		Scanner sc=new Scanner(System.in);
		System.out.println("ENTER NUMBER....");
	
		String s1=	sc.nextLine();
		int [] arraynumber=new int[s1.length()];
		int[] repeatedarray=new int[s1.length()];
		
		for(int i=0;i<s1.length();i++)
		{
			arraynumber[i]=Integer.valueOf(s1.charAt(i))-'0';
		}
		System.out.println("array of number is..."+Arrays.toString(arraynumber));
		
		int countahead=arraynumber.length/2;
		for(int j=0,k=0;j<arraynumber.length && (j+countahead)<arraynumber.length;j++,k=k+2)
		{
			if(arraynumber[j]==arraynumber[j+countahead])
			{
				repeatedarray[k]=arraynumber[j];
				repeatedarray[k+1]=arraynumber[j];
				System.out.println("repeated array is..."+repeatedarray[k]+"   "+repeatedarray[k+1]);
			}
		}
		int sum=0;
		for(int i=0;i<repeatedarray.length;i++)
		{
			sum=sum+repeatedarray[i];
		}
		System.out.println("sum is..."+sum);
		

	}

}
