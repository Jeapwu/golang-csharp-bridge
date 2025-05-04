#ifndef BRIDGE_H
#define BRIDGE_H

#ifdef _WIN32
#include <windows.h>
typedef HINSTANCE LibHandle;
typedef FARPROC SymbolHandle;
#else
#include <stdlib.h>
#include <dlfcn.h>
typedef void* LibHandle;
typedef void* SymbolHandle;
#endif

typedef int (*add_func)(int, int);

static inline LibHandle load_library(const char* path) {
#ifdef _WIN32
    return LoadLibraryA(path);
#else
    return dlopen(path, RTLD_LAZY);
#endif
}

static inline SymbolHandle load_symbol(LibHandle handle, const char* name) {
#ifdef _WIN32
    return GetProcAddress(handle, name);
#else
    return dlsym(handle, name);
#endif
}

static inline void close_library(LibHandle handle) {
#ifdef _WIN32
    if (handle) FreeLibrary(handle);
#else
    if (handle) dlclose(handle);
#endif
}

// C 端中转函数：传函数指针和参数，C 帮你调用
static inline int call_add(add_func fn, int a, int b) {
    return fn(a, b);
}

#endif // BRIDGE_H