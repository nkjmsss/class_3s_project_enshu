#pragma once

#ifdef _DEBUG
#   define MyOutputDebugString( str, ... ) \
      { \
        WCHAR c[256]; \
        swprintf( c, str, __VA_ARGS__ ); \
        OutputDebugString( c ); \
      }
#else
#    define MyOutputDebugString( str, ... ) // ‹óŽÀ‘•
#endif