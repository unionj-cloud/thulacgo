#ifndef THULAC_GO_H
#define THULAC_GO_H

#ifdef __cplusplus
extern "C" {
#endif
typedef void* Thulac;
Thulac NewThulac(const char * model_path, const char* user_path, int just_seg, int t2s, int ufilter, char separator);
void Deinit(Thulac);
const char* Seg(Thulac, const char *in);
void SegToSlice(Thulac l, const char *in, char ***out, int *size);

#ifdef __cplusplus
}
#endif

#endif // THULAC_GO_H
