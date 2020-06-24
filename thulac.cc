#include "thulac.h"
#include "cppthulac/thulac.h"
#include <iostream>
using std::cout;
using std::endl;

Thulac NewThulac(const char * model_path, const char* user_path, int just_seg, int t2s, int ufilter, char separator)
{
    THULAC* lac = new THULAC();
    lac->init(model_path, user_path, just_seg, t2s, ufilter, separator);
    return (Thulac)lac;
}
void Deinit(Thulac l)
{
    ((THULAC*)l)->deinit();
}

const char* Seg(Thulac l, const char *in)
{
    std::string str(in);
    THULAC * lac = (THULAC*)l;
    std::string tostr = lac->toString(lac->cut(str));
    char * cstr = new char [tostr.length()+1];
    std::strcpy (cstr, tostr.c_str());
    return cstr;
}

char *convert(const std::string & s)
{
   char *pc = new char[s.size()+1];
   std::strcpy(pc, s.c_str());
   return pc;
}

void SegToSlice(Thulac l, const char *in, char ***out, int *size)
{
    std::string str(in);
    THULAC * lac = (THULAC*)l;

    std::vector<std::string> vs = lac->toArray(lac->cut(str));
    int vsize = vs.size();

    char ** vc = new char *[vsize];

    for(int i = 0; i < vs.size(); ++i)
    {
        std::string str = vs[i];
        char * cstr = new char [str.length()+1];
        std::strcpy (cstr, str.c_str());
        vc[i] = cstr;
    }

    *out = &vc[0];
    *size = vsize;
}
