import System;

namespace Examples
{
    public interface IWeapon
    {
        string Fire();
    }

    public class Gun : IWeapon
    {
        public string Fire()
        {
            return "Bang!";
        }
    }

    public class Lasergun : IWeapon
    {
        public string Fire()
        {
            return "Pewpew!";
        }
    }

    public class Shotgun : IWeapon
    {
        public string Fire()
        {
            return "Boom!";
        }
    }

    public class Program
    {
        public static void Main(string[] args)
        {
            Console.WriteLine(new Shotgun().Fire()); // Boom!
        }
    }
}