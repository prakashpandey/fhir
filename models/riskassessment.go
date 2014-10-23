// Copyright (c) 2011-2014, HL7, Inc & The MITRE Corporation
// All rights reserved.
//
// Redistribution and use in source and binary forms, with or without modification,
// are permitted provided that the following conditions are met:
//
//     * Redistributions of source code must retain the above copyright notice, this
//       list of conditions and the following disclaimer.
//     * Redistributions in binary form must reproduce the above copyright notice,
//       this list of conditions and the following disclaimer in the documentation
//       and/or other materials provided with the distribution.
//     * Neither the name of HL7 nor the names of its contributors may be used to
//       endorse or promote products derived from this software without specific
//       prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS" AND
// ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE IMPLIED
// WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE DISCLAIMED.
// IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT,
// INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT
// NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR
// PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY,
// WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
// ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
// POSSIBILITY OF SUCH DAMAGE.

package models

type RiskAssessment struct {
	Id         string                              `json:"-" bson:"_id"`
	Subject    Reference                           `bson:"subject"`
	Date       FHIRDateTime                        `bson:"date"`
	Condition  Reference                           `bson:"condition"`
	Performer  Reference                           `bson:"performer"`
	Identifier Identifier                          `bson:"identifier"`
	Method     CodeableConcept                     `bson:"method"`
	Basis      []Reference                         `bson:"basis"`
	Prediction []RiskAssessmentPredictionComponent `bson:"prediction"`
	Mitigation string                              `bson:"mitigation"`
}

// This is an ugly hack to deal with embedded structures in the spec prediction
type RiskAssessmentPredictionComponent struct {
	Outcome                    CodeableConcept `bson:"outcome"`
	ProbabilityDecimal         float64         `bson:"probabilityDecimal"`
	ProbabilityRange           Range           `bson:"probabilityRange"`
	ProbabilityCodeableConcept CodeableConcept `bson:"probabilityCodeableConcept"`
	RelativeRisk               float64         `bson:"relativeRisk"`
	WhenPeriod                 Period          `bson:"whenPeriod"`
	WhenRange                  Range           `bson:"whenRange"`
	Rationale                  string          `bson:"rationale"`
}
