// Package conv provides type conversion helpers.
//
//  - NewT converts T to *T.
//  - NewTFromS converts S to *T.
//  - ValueT converts *T to T, returns zero value in case of nil.
//  - MaybeT converts between protobuf WKT and optional value of corresponding Go type.
//  - MaybeTFromS converts optional S to optional T (optional usually means ref).
//  - Must* is panicking versions of above functions.
package conv
