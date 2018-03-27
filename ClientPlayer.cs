using System;
using System.Net.Sockets;
using System.IO;
public class ClientPlayer
{
	public static void Main(String []args)
	{
		try
		{
			TcpClient socket = new TcpClient("localhost",3000);
			NetworkStream networkStream=socket.GetStream();
			StreamReader reader = new StreamReader(networkStream);
			StreamWriter writer = new StreamWriter(networkStream);
			while(true)
			{
				string readFromServer = reader.ReadLine();
				Console.WriteLine(readFromServer);
			}
			
		}
		catch(Exception e)
		{
			Console.WriteLine(e);
			return;
		}
	}
}