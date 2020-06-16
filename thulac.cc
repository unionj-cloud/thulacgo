#include "thulac.h"
#include "cppthulac/thulac.h"

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
    return tostr.c_str();
}
