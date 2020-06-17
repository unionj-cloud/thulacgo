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

//std::string toString(const THULAC_result& result, bool seg_only, char separator) {
//    std::ostringstream ous;
//
//    for(int i = 0; i < result.size() - 1; i++) {
//        if(i != 0) ous << " ";
//        if(seg_only) {
//            ous << result[i].first;
//        }
//        else {
//            ous << result[i].first << separator << result[i].second;
//        }
//    }
//
//    return ous.str();
//}

const char* Seg(Thulac l, const char *in)
{
    std::string str(in);
    THULAC * lac = (THULAC*)l;
    std::string tostr = lac->toString(lac->cut(str));
    return tostr.c_str();
}
