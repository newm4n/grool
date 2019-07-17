// Generated from /Users/ferdinandneman/Laboratory/Golang/src/github.com/newm4n/grool/antlr/grool.g4 by ANTLR 4.7.1
import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class groolParser extends Parser {
	static { RuntimeMetaData.checkVersion("4.7.1", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		T__0=1, RULE=2, WHEN=3, THEN=4, AND=5, OR=6, TRUE=7, FALSE=8, NULL_LITERAL=9, 
		NOT=10, SIMPLENAME=11, DOTTEDNAME=12, PLUS=13, MINUS=14, DIV=15, MUL=16, 
		EQUALS=17, ASSIGN=18, GT=19, LT=20, GTE=21, LTE=22, NOTEQUALS=23, SEMICOLON=24, 
		LR_BRACE=25, RR_BRACE=26, LR_BRACKET=27, RR_BRACKET=28, DOT=29, DQUOTA_STRING=30, 
		SQUOTA_STRING=31, DECIMAL_LITERAL=32, REAL_LITERAL=33, SPACE=34, COMMENT=35, 
		LINE_COMMENT=36;
	public static final int
		RULE_root = 0, RULE_ruleEntry = 1, RULE_ruleName = 2, RULE_ruleDescription = 3, 
		RULE_whenScope = 4, RULE_thenScope = 5, RULE_assignExpressions = 6, RULE_assignExpression = 7, 
		RULE_assignment = 8, RULE_expression = 9, RULE_predicate = 10, RULE_expressionAtom = 11, 
		RULE_functionCall = 12, RULE_functionArgs = 13, RULE_logicalOperator = 14, 
		RULE_variable = 15, RULE_mathOperator = 16, RULE_comparisonOperator = 17, 
		RULE_constant = 18, RULE_decimalLiteral = 19, RULE_stringLiteral = 20, 
		RULE_booleanLiteral = 21;
	public static final String[] ruleNames = {
		"root", "ruleEntry", "ruleName", "ruleDescription", "whenScope", "thenScope", 
		"assignExpressions", "assignExpression", "assignment", "expression", "predicate", 
		"expressionAtom", "functionCall", "functionArgs", "logicalOperator", "variable", 
		"mathOperator", "comparisonOperator", "constant", "decimalLiteral", "stringLiteral", 
		"booleanLiteral"
	};

	private static final String[] _LITERAL_NAMES = {
		null, "','", null, null, null, "'&&'", "'||'", null, null, null, null, 
		null, null, "'+'", "'-'", "'/'", "'*'", "'=='", "'='", "'>'", "'<'", "'>='", 
		"'<='", "'!='", "';'", "'{'", "'}'", "'('", "')'", "'.'"
	};
	private static final String[] _SYMBOLIC_NAMES = {
		null, null, "RULE", "WHEN", "THEN", "AND", "OR", "TRUE", "FALSE", "NULL_LITERAL", 
		"NOT", "SIMPLENAME", "DOTTEDNAME", "PLUS", "MINUS", "DIV", "MUL", "EQUALS", 
		"ASSIGN", "GT", "LT", "GTE", "LTE", "NOTEQUALS", "SEMICOLON", "LR_BRACE", 
		"RR_BRACE", "LR_BRACKET", "RR_BRACKET", "DOT", "DQUOTA_STRING", "SQUOTA_STRING", 
		"DECIMAL_LITERAL", "REAL_LITERAL", "SPACE", "COMMENT", "LINE_COMMENT"
	};
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "grool.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public groolParser(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}
	public static class RootContext extends ParserRuleContext {
		public TerminalNode EOF() { return getToken(groolParser.EOF, 0); }
		public List<RuleEntryContext> ruleEntry() {
			return getRuleContexts(RuleEntryContext.class);
		}
		public RuleEntryContext ruleEntry(int i) {
			return getRuleContext(RuleEntryContext.class,i);
		}
		public RootContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_root; }
	}

	public final RootContext root() throws RecognitionException {
		RootContext _localctx = new RootContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_root);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(47);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==RULE) {
				{
				{
				setState(44);
				ruleEntry();
				}
				}
				setState(49);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			setState(50);
			match(EOF);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RuleEntryContext extends ParserRuleContext {
		public TerminalNode RULE() { return getToken(groolParser.RULE, 0); }
		public RuleNameContext ruleName() {
			return getRuleContext(RuleNameContext.class,0);
		}
		public TerminalNode LR_BRACE() { return getToken(groolParser.LR_BRACE, 0); }
		public WhenScopeContext whenScope() {
			return getRuleContext(WhenScopeContext.class,0);
		}
		public ThenScopeContext thenScope() {
			return getRuleContext(ThenScopeContext.class,0);
		}
		public TerminalNode RR_BRACE() { return getToken(groolParser.RR_BRACE, 0); }
		public RuleDescriptionContext ruleDescription() {
			return getRuleContext(RuleDescriptionContext.class,0);
		}
		public RuleEntryContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleEntry; }
	}

	public final RuleEntryContext ruleEntry() throws RecognitionException {
		RuleEntryContext _localctx = new RuleEntryContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_ruleEntry);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(52);
			match(RULE);
			setState(53);
			ruleName();
			setState(55);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if (_la==DQUOTA_STRING || _la==SQUOTA_STRING) {
				{
				setState(54);
				ruleDescription();
				}
			}

			setState(57);
			match(LR_BRACE);
			setState(58);
			whenScope();
			setState(59);
			thenScope();
			setState(60);
			match(RR_BRACE);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RuleNameContext extends ParserRuleContext {
		public TerminalNode SIMPLENAME() { return getToken(groolParser.SIMPLENAME, 0); }
		public RuleNameContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleName; }
	}

	public final RuleNameContext ruleName() throws RecognitionException {
		RuleNameContext _localctx = new RuleNameContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_ruleName);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(62);
			match(SIMPLENAME);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RuleDescriptionContext extends ParserRuleContext {
		public TerminalNode DQUOTA_STRING() { return getToken(groolParser.DQUOTA_STRING, 0); }
		public TerminalNode SQUOTA_STRING() { return getToken(groolParser.SQUOTA_STRING, 0); }
		public RuleDescriptionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ruleDescription; }
	}

	public final RuleDescriptionContext ruleDescription() throws RecognitionException {
		RuleDescriptionContext _localctx = new RuleDescriptionContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_ruleDescription);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(64);
			_la = _input.LA(1);
			if ( !(_la==DQUOTA_STRING || _la==SQUOTA_STRING) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class WhenScopeContext extends ParserRuleContext {
		public TerminalNode WHEN() { return getToken(groolParser.WHEN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public WhenScopeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_whenScope; }
	}

	public final WhenScopeContext whenScope() throws RecognitionException {
		WhenScopeContext _localctx = new WhenScopeContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_whenScope);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(66);
			match(WHEN);
			setState(67);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ThenScopeContext extends ParserRuleContext {
		public TerminalNode THEN() { return getToken(groolParser.THEN, 0); }
		public AssignExpressionsContext assignExpressions() {
			return getRuleContext(AssignExpressionsContext.class,0);
		}
		public ThenScopeContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_thenScope; }
	}

	public final ThenScopeContext thenScope() throws RecognitionException {
		ThenScopeContext _localctx = new ThenScopeContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_thenScope);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(69);
			match(THEN);
			setState(70);
			assignExpressions();
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignExpressionsContext extends ParserRuleContext {
		public List<AssignExpressionContext> assignExpression() {
			return getRuleContexts(AssignExpressionContext.class);
		}
		public AssignExpressionContext assignExpression(int i) {
			return getRuleContext(AssignExpressionContext.class,i);
		}
		public AssignExpressionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignExpressions; }
	}

	public final AssignExpressionsContext assignExpressions() throws RecognitionException {
		AssignExpressionsContext _localctx = new AssignExpressionsContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_assignExpressions);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(73); 
			_errHandler.sync(this);
			_la = _input.LA(1);
			do {
				{
				{
				setState(72);
				assignExpression();
				}
				}
				setState(75); 
				_errHandler.sync(this);
				_la = _input.LA(1);
			} while ( _la==SIMPLENAME || _la==DOTTEDNAME );
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignExpressionContext extends ParserRuleContext {
		public AssignmentContext assignment() {
			return getRuleContext(AssignmentContext.class,0);
		}
		public TerminalNode SEMICOLON() { return getToken(groolParser.SEMICOLON, 0); }
		public AssignExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignExpression; }
	}

	public final AssignExpressionContext assignExpression() throws RecognitionException {
		AssignExpressionContext _localctx = new AssignExpressionContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_assignExpression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(77);
			assignment();
			setState(78);
			match(SEMICOLON);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class AssignmentContext extends ParserRuleContext {
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public TerminalNode ASSIGN() { return getToken(groolParser.ASSIGN, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public AssignmentContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_assignment; }
	}

	public final AssignmentContext assignment() throws RecognitionException {
		AssignmentContext _localctx = new AssignmentContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_assignment);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(80);
			variable();
			setState(81);
			match(ASSIGN);
			setState(82);
			expression(0);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public TerminalNode LR_BRACKET() { return getToken(groolParser.LR_BRACKET, 0); }
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public LogicalOperatorContext logicalOperator() {
			return getRuleContext(LogicalOperatorContext.class,0);
		}
		public TerminalNode RR_BRACKET() { return getToken(groolParser.RR_BRACKET, 0); }
		public PredicateContext predicate() {
			return getRuleContext(PredicateContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		return expression(0);
	}

	private ExpressionContext expression(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionContext _localctx = new ExpressionContext(_ctx, _parentState);
		ExpressionContext _prevctx = _localctx;
		int _startState = 18;
		enterRecursionRule(_localctx, 18, RULE_expression, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(92);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case LR_BRACKET:
				{
				setState(85);
				match(LR_BRACKET);
				setState(86);
				expression(0);
				setState(87);
				logicalOperator();
				setState(88);
				expression(0);
				setState(89);
				match(RR_BRACKET);
				}
				break;
			case TRUE:
			case FALSE:
			case NULL_LITERAL:
			case NOT:
			case SIMPLENAME:
			case DOTTEDNAME:
			case MINUS:
			case DQUOTA_STRING:
			case SQUOTA_STRING:
			case DECIMAL_LITERAL:
			case REAL_LITERAL:
				{
				setState(91);
				predicate();
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(100);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionContext(_parentctx, _parentState);
					pushNewRecursionContext(_localctx, _startState, RULE_expression);
					setState(94);
					if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
					setState(95);
					logicalOperator();
					setState(96);
					expression(4);
					}
					} 
				}
				setState(102);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,4,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PredicateContext extends ParserRuleContext {
		public List<ExpressionAtomContext> expressionAtom() {
			return getRuleContexts(ExpressionAtomContext.class);
		}
		public ExpressionAtomContext expressionAtom(int i) {
			return getRuleContext(ExpressionAtomContext.class,i);
		}
		public ComparisonOperatorContext comparisonOperator() {
			return getRuleContext(ComparisonOperatorContext.class,0);
		}
		public PredicateContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_predicate; }
	}

	public final PredicateContext predicate() throws RecognitionException {
		PredicateContext _localctx = new PredicateContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_predicate);
		try {
			setState(108);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,5,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(103);
				expressionAtom(0);
				setState(104);
				comparisonOperator();
				setState(105);
				expressionAtom(0);
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(107);
				expressionAtom(0);
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionAtomContext extends ParserRuleContext {
		public ExpressionAtomContext left;
		public ExpressionAtomContext right;
		public ConstantContext constant() {
			return getRuleContext(ConstantContext.class,0);
		}
		public VariableContext variable() {
			return getRuleContext(VariableContext.class,0);
		}
		public FunctionCallContext functionCall() {
			return getRuleContext(FunctionCallContext.class,0);
		}
		public MathOperatorContext mathOperator() {
			return getRuleContext(MathOperatorContext.class,0);
		}
		public List<ExpressionAtomContext> expressionAtom() {
			return getRuleContexts(ExpressionAtomContext.class);
		}
		public ExpressionAtomContext expressionAtom(int i) {
			return getRuleContext(ExpressionAtomContext.class,i);
		}
		public ExpressionAtomContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expressionAtom; }
	}

	public final ExpressionAtomContext expressionAtom() throws RecognitionException {
		return expressionAtom(0);
	}

	private ExpressionAtomContext expressionAtom(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ExpressionAtomContext _localctx = new ExpressionAtomContext(_ctx, _parentState);
		ExpressionAtomContext _prevctx = _localctx;
		int _startState = 22;
		enterRecursionRule(_localctx, 22, RULE_expressionAtom, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(114);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				{
				setState(111);
				constant();
				}
				break;
			case 2:
				{
				setState(112);
				variable();
				}
				break;
			case 3:
				{
				setState(113);
				functionCall();
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(122);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ExpressionAtomContext(_parentctx, _parentState);
					_localctx.left = _prevctx;
					_localctx.left = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_expressionAtom);
					setState(116);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(117);
					mathOperator();
					setState(118);
					((ExpressionAtomContext)_localctx).right = expressionAtom(3);
					}
					} 
				}
				setState(124);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,7,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class FunctionCallContext extends ParserRuleContext {
		public TerminalNode SIMPLENAME() { return getToken(groolParser.SIMPLENAME, 0); }
		public FunctionArgsContext functionArgs() {
			return getRuleContext(FunctionArgsContext.class,0);
		}
		public FunctionCallContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionCall; }
	}

	public final FunctionCallContext functionCall() throws RecognitionException {
		FunctionCallContext _localctx = new FunctionCallContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_functionCall);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(125);
			match(SIMPLENAME);
			setState(126);
			match(LR_BRACKET);
			setState(128);
			_errHandler.sync(this);
			_la = _input.LA(1);
			if ((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << TRUE) | (1L << FALSE) | (1L << NULL_LITERAL) | (1L << NOT) | (1L << SIMPLENAME) | (1L << DOTTEDNAME) | (1L << MINUS) | (1L << LR_BRACKET) | (1L << DQUOTA_STRING) | (1L << SQUOTA_STRING) | (1L << DECIMAL_LITERAL) | (1L << REAL_LITERAL))) != 0)) {
				{
				setState(127);
				functionArgs();
				}
			}

			setState(130);
			match(RR_BRACKET);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FunctionArgsContext extends ParserRuleContext {
		public List<ConstantContext> constant() {
			return getRuleContexts(ConstantContext.class);
		}
		public ConstantContext constant(int i) {
			return getRuleContext(ConstantContext.class,i);
		}
		public List<VariableContext> variable() {
			return getRuleContexts(VariableContext.class);
		}
		public VariableContext variable(int i) {
			return getRuleContext(VariableContext.class,i);
		}
		public List<FunctionCallContext> functionCall() {
			return getRuleContexts(FunctionCallContext.class);
		}
		public FunctionCallContext functionCall(int i) {
			return getRuleContext(FunctionCallContext.class,i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public FunctionArgsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_functionArgs; }
	}

	public final FunctionArgsContext functionArgs() throws RecognitionException {
		FunctionArgsContext _localctx = new FunctionArgsContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_functionArgs);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(136);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
			case 1:
				{
				setState(132);
				constant();
				}
				break;
			case 2:
				{
				setState(133);
				variable();
				}
				break;
			case 3:
				{
				setState(134);
				functionCall();
				}
				break;
			case 4:
				{
				setState(135);
				expression(0);
				}
				break;
			}
			setState(147);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (_la==T__0) {
				{
				{
				setState(138);
				match(T__0);
				setState(143);
				_errHandler.sync(this);
				switch ( getInterpreter().adaptivePredict(_input,10,_ctx) ) {
				case 1:
					{
					setState(139);
					constant();
					}
					break;
				case 2:
					{
					setState(140);
					variable();
					}
					break;
				case 3:
					{
					setState(141);
					functionCall();
					}
					break;
				case 4:
					{
					setState(142);
					expression(0);
					}
					break;
				}
				}
				}
				setState(149);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LogicalOperatorContext extends ParserRuleContext {
		public TerminalNode AND() { return getToken(groolParser.AND, 0); }
		public TerminalNode OR() { return getToken(groolParser.OR, 0); }
		public LogicalOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_logicalOperator; }
	}

	public final LogicalOperatorContext logicalOperator() throws RecognitionException {
		LogicalOperatorContext _localctx = new LogicalOperatorContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_logicalOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(150);
			_la = _input.LA(1);
			if ( !(_la==AND || _la==OR) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VariableContext extends ParserRuleContext {
		public TerminalNode SIMPLENAME() { return getToken(groolParser.SIMPLENAME, 0); }
		public TerminalNode DOTTEDNAME() { return getToken(groolParser.DOTTEDNAME, 0); }
		public VariableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_variable; }
	}

	public final VariableContext variable() throws RecognitionException {
		VariableContext _localctx = new VariableContext(_ctx, getState());
		enterRule(_localctx, 30, RULE_variable);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(152);
			_la = _input.LA(1);
			if ( !(_la==SIMPLENAME || _la==DOTTEDNAME) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MathOperatorContext extends ParserRuleContext {
		public TerminalNode MUL() { return getToken(groolParser.MUL, 0); }
		public TerminalNode DIV() { return getToken(groolParser.DIV, 0); }
		public TerminalNode PLUS() { return getToken(groolParser.PLUS, 0); }
		public TerminalNode MINUS() { return getToken(groolParser.MINUS, 0); }
		public MathOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mathOperator; }
	}

	public final MathOperatorContext mathOperator() throws RecognitionException {
		MathOperatorContext _localctx = new MathOperatorContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_mathOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(154);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << PLUS) | (1L << MINUS) | (1L << DIV) | (1L << MUL))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ComparisonOperatorContext extends ParserRuleContext {
		public TerminalNode GT() { return getToken(groolParser.GT, 0); }
		public TerminalNode LT() { return getToken(groolParser.LT, 0); }
		public TerminalNode GTE() { return getToken(groolParser.GTE, 0); }
		public TerminalNode LTE() { return getToken(groolParser.LTE, 0); }
		public TerminalNode EQUALS() { return getToken(groolParser.EQUALS, 0); }
		public TerminalNode NOTEQUALS() { return getToken(groolParser.NOTEQUALS, 0); }
		public ComparisonOperatorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_comparisonOperator; }
	}

	public final ComparisonOperatorContext comparisonOperator() throws RecognitionException {
		ComparisonOperatorContext _localctx = new ComparisonOperatorContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_comparisonOperator);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(156);
			_la = _input.LA(1);
			if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << EQUALS) | (1L << GT) | (1L << LT) | (1L << GTE) | (1L << LTE) | (1L << NOTEQUALS))) != 0)) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConstantContext extends ParserRuleContext {
		public StringLiteralContext stringLiteral() {
			return getRuleContext(StringLiteralContext.class,0);
		}
		public DecimalLiteralContext decimalLiteral() {
			return getRuleContext(DecimalLiteralContext.class,0);
		}
		public BooleanLiteralContext booleanLiteral() {
			return getRuleContext(BooleanLiteralContext.class,0);
		}
		public TerminalNode REAL_LITERAL() { return getToken(groolParser.REAL_LITERAL, 0); }
		public TerminalNode NULL_LITERAL() { return getToken(groolParser.NULL_LITERAL, 0); }
		public TerminalNode NOT() { return getToken(groolParser.NOT, 0); }
		public ConstantContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_constant; }
	}

	public final ConstantContext constant() throws RecognitionException {
		ConstantContext _localctx = new ConstantContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_constant);
		int _la;
		try {
			setState(168);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DQUOTA_STRING:
			case SQUOTA_STRING:
				enterOuterAlt(_localctx, 1);
				{
				setState(158);
				stringLiteral();
				}
				break;
			case DECIMAL_LITERAL:
				enterOuterAlt(_localctx, 2);
				{
				setState(159);
				decimalLiteral();
				}
				break;
			case MINUS:
				enterOuterAlt(_localctx, 3);
				{
				setState(160);
				match(MINUS);
				setState(161);
				decimalLiteral();
				}
				break;
			case TRUE:
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(162);
				booleanLiteral();
				}
				break;
			case REAL_LITERAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(163);
				match(REAL_LITERAL);
				}
				break;
			case NULL_LITERAL:
			case NOT:
				enterOuterAlt(_localctx, 6);
				{
				setState(165);
				_errHandler.sync(this);
				_la = _input.LA(1);
				if (_la==NOT) {
					{
					setState(164);
					match(NOT);
					}
				}

				setState(167);
				match(NULL_LITERAL);
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DecimalLiteralContext extends ParserRuleContext {
		public TerminalNode DECIMAL_LITERAL() { return getToken(groolParser.DECIMAL_LITERAL, 0); }
		public DecimalLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_decimalLiteral; }
	}

	public final DecimalLiteralContext decimalLiteral() throws RecognitionException {
		DecimalLiteralContext _localctx = new DecimalLiteralContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_decimalLiteral);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			match(DECIMAL_LITERAL);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class StringLiteralContext extends ParserRuleContext {
		public TerminalNode DQUOTA_STRING() { return getToken(groolParser.DQUOTA_STRING, 0); }
		public TerminalNode SQUOTA_STRING() { return getToken(groolParser.SQUOTA_STRING, 0); }
		public StringLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_stringLiteral; }
	}

	public final StringLiteralContext stringLiteral() throws RecognitionException {
		StringLiteralContext _localctx = new StringLiteralContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_stringLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(172);
			_la = _input.LA(1);
			if ( !(_la==DQUOTA_STRING || _la==SQUOTA_STRING) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BooleanLiteralContext extends ParserRuleContext {
		public TerminalNode TRUE() { return getToken(groolParser.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(groolParser.FALSE, 0); }
		public BooleanLiteralContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_booleanLiteral; }
	}

	public final BooleanLiteralContext booleanLiteral() throws RecognitionException {
		BooleanLiteralContext _localctx = new BooleanLiteralContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_booleanLiteral);
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(174);
			_la = _input.LA(1);
			if ( !(_la==TRUE || _la==FALSE) ) {
			_errHandler.recoverInline(this);
			}
			else {
				if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
				_errHandler.reportMatch(this);
				consume();
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 9:
			return expression_sempred((ExpressionContext)_localctx, predIndex);
		case 11:
			return expressionAtom_sempred((ExpressionAtomContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean expression_sempred(ExpressionContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean expressionAtom_sempred(ExpressionAtomContext _localctx, int predIndex) {
		switch (predIndex) {
		case 1:
			return precpred(_ctx, 2);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3&\u00b3\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\3\2\7\2\60\n\2\f\2"+
		"\16\2\63\13\2\3\2\3\2\3\3\3\3\3\3\5\3:\n\3\3\3\3\3\3\3\3\3\3\3\3\4\3\4"+
		"\3\5\3\5\3\6\3\6\3\6\3\7\3\7\3\7\3\b\6\bL\n\b\r\b\16\bM\3\t\3\t\3\t\3"+
		"\n\3\n\3\n\3\n\3\13\3\13\3\13\3\13\3\13\3\13\3\13\3\13\5\13_\n\13\3\13"+
		"\3\13\3\13\3\13\7\13e\n\13\f\13\16\13h\13\13\3\f\3\f\3\f\3\f\3\f\5\fo"+
		"\n\f\3\r\3\r\3\r\3\r\5\ru\n\r\3\r\3\r\3\r\3\r\7\r{\n\r\f\r\16\r~\13\r"+
		"\3\16\3\16\3\16\5\16\u0083\n\16\3\16\3\16\3\17\3\17\3\17\3\17\5\17\u008b"+
		"\n\17\3\17\3\17\3\17\3\17\3\17\5\17\u0092\n\17\7\17\u0094\n\17\f\17\16"+
		"\17\u0097\13\17\3\20\3\20\3\21\3\21\3\22\3\22\3\23\3\23\3\24\3\24\3\24"+
		"\3\24\3\24\3\24\3\24\5\24\u00a8\n\24\3\24\5\24\u00ab\n\24\3\25\3\25\3"+
		"\26\3\26\3\27\3\27\3\27\2\4\24\30\30\2\4\6\b\n\f\16\20\22\24\26\30\32"+
		"\34\36 \"$&(*,\2\b\3\2 !\3\2\7\b\3\2\r\16\3\2\17\22\4\2\23\23\25\31\3"+
		"\2\t\n\2\u00b3\2\61\3\2\2\2\4\66\3\2\2\2\6@\3\2\2\2\bB\3\2\2\2\nD\3\2"+
		"\2\2\fG\3\2\2\2\16K\3\2\2\2\20O\3\2\2\2\22R\3\2\2\2\24^\3\2\2\2\26n\3"+
		"\2\2\2\30t\3\2\2\2\32\177\3\2\2\2\34\u008a\3\2\2\2\36\u0098\3\2\2\2 \u009a"+
		"\3\2\2\2\"\u009c\3\2\2\2$\u009e\3\2\2\2&\u00aa\3\2\2\2(\u00ac\3\2\2\2"+
		"*\u00ae\3\2\2\2,\u00b0\3\2\2\2.\60\5\4\3\2/.\3\2\2\2\60\63\3\2\2\2\61"+
		"/\3\2\2\2\61\62\3\2\2\2\62\64\3\2\2\2\63\61\3\2\2\2\64\65\7\2\2\3\65\3"+
		"\3\2\2\2\66\67\7\4\2\2\679\5\6\4\28:\5\b\5\298\3\2\2\29:\3\2\2\2:;\3\2"+
		"\2\2;<\7\33\2\2<=\5\n\6\2=>\5\f\7\2>?\7\34\2\2?\5\3\2\2\2@A\7\r\2\2A\7"+
		"\3\2\2\2BC\t\2\2\2C\t\3\2\2\2DE\7\5\2\2EF\5\24\13\2F\13\3\2\2\2GH\7\6"+
		"\2\2HI\5\16\b\2I\r\3\2\2\2JL\5\20\t\2KJ\3\2\2\2LM\3\2\2\2MK\3\2\2\2MN"+
		"\3\2\2\2N\17\3\2\2\2OP\5\22\n\2PQ\7\32\2\2Q\21\3\2\2\2RS\5 \21\2ST\7\24"+
		"\2\2TU\5\24\13\2U\23\3\2\2\2VW\b\13\1\2WX\7\35\2\2XY\5\24\13\2YZ\5\36"+
		"\20\2Z[\5\24\13\2[\\\7\36\2\2\\_\3\2\2\2]_\5\26\f\2^V\3\2\2\2^]\3\2\2"+
		"\2_f\3\2\2\2`a\f\5\2\2ab\5\36\20\2bc\5\24\13\6ce\3\2\2\2d`\3\2\2\2eh\3"+
		"\2\2\2fd\3\2\2\2fg\3\2\2\2g\25\3\2\2\2hf\3\2\2\2ij\5\30\r\2jk\5$\23\2"+
		"kl\5\30\r\2lo\3\2\2\2mo\5\30\r\2ni\3\2\2\2nm\3\2\2\2o\27\3\2\2\2pq\b\r"+
		"\1\2qu\5&\24\2ru\5 \21\2su\5\32\16\2tp\3\2\2\2tr\3\2\2\2ts\3\2\2\2u|\3"+
		"\2\2\2vw\f\4\2\2wx\5\"\22\2xy\5\30\r\5y{\3\2\2\2zv\3\2\2\2{~\3\2\2\2|"+
		"z\3\2\2\2|}\3\2\2\2}\31\3\2\2\2~|\3\2\2\2\177\u0080\7\r\2\2\u0080\u0082"+
		"\7\35\2\2\u0081\u0083\5\34\17\2\u0082\u0081\3\2\2\2\u0082\u0083\3\2\2"+
		"\2\u0083\u0084\3\2\2\2\u0084\u0085\7\36\2\2\u0085\33\3\2\2\2\u0086\u008b"+
		"\5&\24\2\u0087\u008b\5 \21\2\u0088\u008b\5\32\16\2\u0089\u008b\5\24\13"+
		"\2\u008a\u0086\3\2\2\2\u008a\u0087\3\2\2\2\u008a\u0088\3\2\2\2\u008a\u0089"+
		"\3\2\2\2\u008b\u0095\3\2\2\2\u008c\u0091\7\3\2\2\u008d\u0092\5&\24\2\u008e"+
		"\u0092\5 \21\2\u008f\u0092\5\32\16\2\u0090\u0092\5\24\13\2\u0091\u008d"+
		"\3\2\2\2\u0091\u008e\3\2\2\2\u0091\u008f\3\2\2\2\u0091\u0090\3\2\2\2\u0092"+
		"\u0094\3\2\2\2\u0093\u008c\3\2\2\2\u0094\u0097\3\2\2\2\u0095\u0093\3\2"+
		"\2\2\u0095\u0096\3\2\2\2\u0096\35\3\2\2\2\u0097\u0095\3\2\2\2\u0098\u0099"+
		"\t\3\2\2\u0099\37\3\2\2\2\u009a\u009b\t\4\2\2\u009b!\3\2\2\2\u009c\u009d"+
		"\t\5\2\2\u009d#\3\2\2\2\u009e\u009f\t\6\2\2\u009f%\3\2\2\2\u00a0\u00ab"+
		"\5*\26\2\u00a1\u00ab\5(\25\2\u00a2\u00a3\7\20\2\2\u00a3\u00ab\5(\25\2"+
		"\u00a4\u00ab\5,\27\2\u00a5\u00ab\7#\2\2\u00a6\u00a8\7\f\2\2\u00a7\u00a6"+
		"\3\2\2\2\u00a7\u00a8\3\2\2\2\u00a8\u00a9\3\2\2\2\u00a9\u00ab\7\13\2\2"+
		"\u00aa\u00a0\3\2\2\2\u00aa\u00a1\3\2\2\2\u00aa\u00a2\3\2\2\2\u00aa\u00a4"+
		"\3\2\2\2\u00aa\u00a5\3\2\2\2\u00aa\u00a7\3\2\2\2\u00ab\'\3\2\2\2\u00ac"+
		"\u00ad\7\"\2\2\u00ad)\3\2\2\2\u00ae\u00af\t\2\2\2\u00af+\3\2\2\2\u00b0"+
		"\u00b1\t\7\2\2\u00b1-\3\2\2\2\20\619M^fnt|\u0082\u008a\u0091\u0095\u00a7"+
		"\u00aa";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}