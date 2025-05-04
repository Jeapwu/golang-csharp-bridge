using System.Runtime.CompilerServices;
using System.Runtime.InteropServices;

public static class NativeExports
{
    [UnmanagedCallersOnly(EntryPoint = "add", CallConvs = new[] { typeof(CallConvCdecl) })]
    public static int Add(int a, int b) => a + b;
}
