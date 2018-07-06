package problemsolving;

import java.util.Scanner;

public class Part1Problem3 {

	public static void main(String[] args) {
		System.out.println("Enter the number to print sprial");
		 Scanner scr=new Scanner(System.in);


		 int number=scr.nextInt();
		 int one=25;
		 //int one=1;
		 ///first step is to find the row and column of value 1 that is in place of one put 1
		 //like int one=1; that row is=303 column=303
		 int oneonerow=2;
		 int oneonecol=2;



		 int array[][]=new int[number][number];
		 int square=number*number;
		  int c1=0;
		  int c2=number-1;
		  int r1=0;
		  int r2=number-1;
		  
		  int roooo=0;
		  int coooo=0;
		  while(square>=1)
		  {
		  for(int i=c1;i<=c2;i++)
		          {
		  
		  array[r1][i]=square--;
		  if(array[r1][i]==one)
		  {
		  System.out.println(r1+"   "+i);
		   roooo=r1;
		   coooo=i;
		  }
		          }

		          for(int j=r1+1;j<=r2;j++)
		          {
		           array[j][c2]=square--;
		           if(array[j][c2]==one)
		  {
		  System.out.println(j+"   "+c2);
		  roooo=j;
		  coooo=c2;
		  }
		          }

		          for(int i=c2-1;i>=c1;i--)
		          {
		           array[r2][i]=square--;
		           if(array[r2][i]==one)
		  {
		  System.out.println(r2+"   "+i);
		  roooo=r2;
		   coooo=i;
		  }
		          }

		          for(int j=r2-1;j>=r1+1;j--)
		          {
		           array[j][c1]=square--;
		           if(array[j][c1]==one)
		  {
		  System.out.println(j+"   "+c1);
		  roooo=j;
		   coooo=c1;
		  }
		          }
		          
		          c1++;
		          c2--;
		          r1++;
		          r2--;
		  }
		  System.out.println("The Circular Matrix is:");
		      for(int i=0;i<number;i++)
		          {
		              for(int j=0;j<number;j++)
		                  {
		                      System.out.print(array[i][j]+ "\t");
		                  }
		           System.out.println();
		          }

		      System.out.println("divyadivyadiyva");
		      System.out.println("roooo    "+roooo);
		      System.out.println("coooo    "+coooo);
		      //System.out.println(array[roooo][]);
		      
		      /*int valueOfOne[]=new int[2];
		      valueOfOne[0]=roooo;
		      valueOfOne[1]=coooo;*/
		      
		     ///row column of 1 respectively=   2  2
		      //row coumn of given number   =   0  371
		      //do subtraction and add sum like (2-0)+(371-2)=371
		      
		      
		      
		      
		      
		      
		      
		 

	}

}
