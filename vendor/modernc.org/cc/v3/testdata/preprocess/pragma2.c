#define __clang__

#if defined(__clang__) && defined(__cplusplus)
#define __MISMATCH_TAGS_PUSH                                            \
        _Pragma("clang diagnostic push")                                \
        _Pragma("clang diagnostic ignored \"-Wmismatched-tags\"")
#define __MISMATCH_TAGS_POP                                             \
        _Pragma("clang diagnostic pop")
#else
#define __MISMATCH_TAGS_PUSH
#define __MISMATCH_TAGS_POP
#endif

#if defined(__clang__)
#define __NULLABILITY_COMPLETENESS_PUSH \
        _Pragma("clang diagnostic push") \
        _Pragma("clang diagnostic ignored \"-Wnullability-completeness\"")
#define __NULLABILITY_COMPLETENESS_POP \
        _Pragma("clang diagnostic pop")
#else
#define __NULLABILITY_COMPLETENESS_PUSH
#define __NULLABILITY_COMPLETENESS_POP
#endif

#define SLIST_HEAD(name, type)                                          \
__MISMATCH_TAGS_PUSH                                                    \
__NULLABILITY_COMPLETENESS_PUSH                                         \
struct name {                                                           \
        struct type *slh_first; /* first element */                     \
}                                                                       \
__NULLABILITY_COMPLETENESS_POP                                          \
__MISMATCH_TAGS_POP

struct knote;
SLIST_HEAD(klist, knote);
